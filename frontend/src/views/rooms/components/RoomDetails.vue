<template>
  <v-toolbar color="#79AE6F" title="Room Details">
    <v-btn v-if="!editMode" icon @click="toggleEdit">
      <v-icon>mdi-pencil</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="cancelEdit">
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="saveRoom" :loading="saving">
      <v-icon>mdi-check</v-icon>
    </v-btn>
  </v-toolbar>

  <v-card>
    <v-card-text>
      <div v-if="!room" class="text-center text-grey">
        Select a room to view details
      </div>
      <v-form ref="roomFormRef" v-else>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="formRoom.roomNo"
              label="Room Number"
              :readonly="!editMode"
              :rules="[rules.required]"
            />
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.number="formRoom.roomWidth"
              label="Width (ft)"
              :readonly="!editMode"
              type="number"
              min="0"
              max="30"
              step="0.25"
              :rules="[rules.positiveNumber]"
            />
          </v-col>

          <v-col cols="12" md="4">
            <v-text-field
              v-model.number="formRoom.roomLength"
              label="Length(ft)"
              :readonly="!editMode"
              type="number"
              min="0"
              max="30"
              step="0.25"
              :rules="[rules.positiveNumber]"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.number="formRoom.doubleBeds"
              label="Double Beds"
              :readonly="!editMode"
              type="number"
              min="0"
              max="10"
              :rules="[rules.positiveNumber]"
            />
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.number="formRoom.singleBeds"
              label="Single Beds"
              type="number"
              :readonly="!editMode"
              min="0"
              max="10"
              :rules="[rules.positiveNumber]"
            />
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model.number="formRoom.windows"
              label="Windows"
              :readonly="!editMode"
              type="number"
              min="0"
              max="10"
              :rules="[rules.positiveNumber]"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="2">
            <v-checkbox label="AC" v-model="formRoom.aC" :readonly="!editMode"></v-checkbox>
          </v-col>
          <v-col cols="12" md="2">
            <v-checkbox label="Wifi" v-model="formRoom.wifi" :readonly="!editMode"></v-checkbox>
          </v-col>
          <v-col cols="12" md="2">
            <v-checkbox label="Hot Water" v-model="formRoom.hotWater" :readonly="!editMode"></v-checkbox>
          </v-col>
          <v-col cols="12" md="2">
            <v-checkbox label="Baclony" v-model="formRoom.balcony" :readonly="!editMode"></v-checkbox>
          </v-col>
          <v-col cols="12" md="4">
            <v-select :items="locations" v-model="formRoom.location" label="Location" :readonly="!editMode"></v-select>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field v-model.number="formRoom.roomCharges" label="Charges" type="number" prefix="$" :rules="[rules.positiveNumber]" :readonly="!editMode"></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-select :items="stati" v-model="formRoom.status" label="Status" :readonly="!editMode"></v-select>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="formRoom.remarks" label="Remarks" :readonly="!editMode"></v-text-field>
          </v-col>
        </v-row>

        <!-- Image upload section (only in edit mode) -->
        <v-row v-if="editMode">
          <v-col cols="12">
            <v-file-input
              ref="fileInputRef"
              label="Room Image"
              accept="image/*"
              prepend-icon="mdi-camera"
              @change="onImageSelected"
              :loading="uploading"
            ></v-file-input>
          </v-col>
        </v-row>

        <!-- Image preview (always visible, even in view mode) -->
        <v-row>
          <v-col cols="12" class="text-center">
            <v-img
              v-if="imageUrl"
              :src="imageUrl"
              max-height="200"
              contain
              class="mb-2"
            ></v-img>
            <div v-else class="text-grey text-caption mb-2">
              No image available
            </div>
            <!-- Optional: delete/clear button (only in edit mode) -->
            <v-btn
              v-if="editMode && formRoom?.roomImage"
              variant="text"
              color="error"
              size="small"
              prepend-icon="mdi-delete"
              @click="clearImage"
            >
              Remove Image
            </v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { Room } from '@/types/index'
import { uploadImage } from '@/services/images'

const props = defineProps<{
  room: Room | null
  locations: string[]
  stati:string[]
}>()

const emit = defineEmits<{
  (e: 'show-message', message: string, color?: string): void
  (e: 'update-room', room: Room): void
}>()

// Local state
const editMode = ref(false)
const saving = ref(false)
const roomFormRef = ref()
const editableRoom = ref<Room | null>(null)   // local copy for editing

const formRoom=computed(()=>editableRoom.value!)

// Validation rules (same as before)
const rules = {
  required: (v: string) => !!v || 'Required',
  positiveNumber: (v: any) => {
    const num = Number(v)
    return (!isNaN(num) && num >= 0)
      || 'Value must be a non-negative number'
  }
}

// When the selected room changes, reset edit mode and update the local copy
watch(() => props.room, (newRoom) => {
  editMode.value = false
  if (newRoom) {
    // Clone the room so we don't mutate the prop
    editableRoom.value = { ...newRoom }
  } else {
    editableRoom.value = null
  }
}, { immediate: true })

function toggleEdit() {
  if (editMode.value) {
    // Cancelling: restore original room (keep the prop's value)
    if (props.room) {
      editableRoom.value = { ...props.room }
    }
    editMode.value = false
  } else {
    // Enter edit mode: ensure we have a fresh copy of the current room
    if (props.room) {
      editableRoom.value = { ...props.room }
    }
    editMode.value = true
  }
}

function cancelEdit() {
  toggleEdit()   // resets form and exits edit mode
}

async function saveRoom() {
  // Validate the form
  const { valid } = await roomFormRef.value?.validate()
  if (!valid) return

  // Emit the local copy (which contains the edits)
  if (editableRoom.value) {
    saving.value = true
    emit('update-room', editableRoom.value)
    saving.value = false
  }
}

// Expose toggleEdit so the parent can manually exit edit mode after a successful save
defineExpose({ toggleEdit })

// Additional refs
const imageFile = ref<File | null>(null)
const uploading = ref(false)
const fileInputRef = ref() // to reset the input after upload

// Function called when a file is selected
async function onImageSelected(event: any) {
  // Vuetify's v-file-input emits the file(s) directly, not an event object
  let file: File | null = null

  if (event instanceof File) {
    file = event
  } else if (Array.isArray(event) && event.length > 0) {
    file = event[0]
  } else if (event?.target?.files) {
    file = event.target.files[0]
  }

  if (!file) {
    console.warn('No file selected')
    return
  }

  uploading.value = true
  try {
    const formData = new FormData()
    formData.append('image', file)

    const response = await uploadImage(formData)
    const imagePath = response.path

    if (formRoom.value) {
      formRoom.value.roomImage = imagePath
    }

    emit('show-message', 'Image uploaded successfully')
  } catch (error) {
    console.error('Upload failed', error)
    emit('show-message', 'Failed to upload image', 'error')
  } finally {
    uploading.value = false
    fileInputRef.value?.reset()
  }
}

// Optional: clear the image (set path to empty)
function clearImage() {
  if (formRoom.value) {
    formRoom.value.roomImage = ''
  }
}

const imageUrl = computed(() => {
  if (!formRoom.value?.roomImage) return null
  // If the path already starts with http, return as is; otherwise prepend backend URL
  if (formRoom.value.roomImage.startsWith('http')) {
    return formRoom.value.roomImage
  }
  // Assuming your backend serves static files at localhost:8080
  return `http://localhost:8080${formRoom.value.roomImage}`
})
</script>