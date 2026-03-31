<template>
  <v-toolbar color="#79AE6F" title="User Details">
    <v-btn v-if="!editMode" icon @click="toggleEdit">
      <v-icon>mdi-pencil</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="cancelEdit">
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="saveUser" :loading="saving">
      <v-icon>mdi-check</v-icon>
    </v-btn>
  </v-toolbar>

  <v-card>
    <v-card-text>
      <div v-if="!user" class="text-center text-grey">
        Select a user to view details
      </div>
      <v-form ref="userFormRef" v-else>
        <v-row>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formUser.userName"
              label="User Name"
              :readonly="!editMode"
              :rules="[rules.required]"
            />
          </v-col>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formUser.password"
              label="Password"
              :readonly="!editMode"
              :rules="[rules.required]"
              type="password"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formUser.userRole"
              label="User Role"
              :readonly="!editMode"
              :rules="[rules.required]"
            />
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { User } from '@/types/index'

const props = defineProps<{
  user: User | null
}>()

const emit = defineEmits<{
  (e: 'show-message', message: string, color?: string): void
  (e: 'update-user', user: User): void
}>()

// Local state
const editMode = ref(false)
const saving = ref(false)
const userFormRef = ref()
const editableUser = ref<User | null>(null) 

const formUser=computed(()=>editableUser.value!)

const rules = {
  required: (v: string) => !!v || 'Required'
}

watch(() => props.user, (newUser) => {
  editMode.value = false
  if (newUser) {
    // Clone the user so we don't mutate the prop
    editableUser.value = { ...newUser }
  } else {
    editableUser.value = null
  }
}, { immediate: true })

function toggleEdit() {
  if (editMode.value) {
    // Cancelling: restore original user (keep the prop's value)
    if (props.user) {
      editableUser.value = { ...props.user }
    }
    editMode.value = false
  } else {
    // Enter edit mode: ensure we have a fresh copy of the current user
    if (props.user) {
      editableUser.value = { ...props.user }
    }
    editMode.value = true
  }
}

function cancelEdit() {
  toggleEdit()   // resets form and exits edit mode
}

async function saveUser() {
  // Validate the form
  const { valid } = await userFormRef.value?.validate()
  if (!valid) return

  // Emit the local copy (which contains the edits)
  if (editableUser.value) {
    saving.value = true
    emit('update-user', editableUser.value)
    saving.value = false
  }
}

// Expose toggleEdit so the parent can manually exit edit mode after a successful save
defineExpose({ toggleEdit })
</script>