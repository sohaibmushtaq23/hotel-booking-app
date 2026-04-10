<template>
  <v-toolbar color="#79AE6F" title="Client Details">
    <v-btn v-if="!editMode" icon @click="toggleEdit">
      <v-icon>mdi-pencil</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="cancelEdit">
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-btn v-if="editMode" icon @click="saveClient" :loading="saving">
      <v-icon>mdi-check</v-icon>
    </v-btn>
  </v-toolbar>

  <v-card>
    <v-card-text>
      <div v-if="!client" class="text-center text-grey">
        Select a client to view details
      </div>
      <v-form ref="clientFormRef" v-else>
        <v-row>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formClient.clientName"
              label="Client Name"
              :readonly="!editMode"
              :rules="[rules.required]"
            />
          </v-col>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formClient.cnic"
              label="CNIC"
              :readonly="!editMode"
              :rules="[rules.required]"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formClient.phone"
              label="Phone"
              :readonly="!editMode"
            />
          </v-col>
          <v-col cols="12" md="6">
            <v-text-field
              v-model="formClient.email"
              label="Email"
              :readonly="!editMode"
              :rules="[rules.email]"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="6">
            <v-text-field
              v-model.number="formClient.discount"
              label="Discount"
              type="number"
              :readonly="!editMode"
              suffix="%"
              step="0.25"
              min="0"
              max="20"
              :rules="[rules.discount]"
            />
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import type { Client } from '@/types/index'

const props = defineProps<{
  client: Client | null
}>()

const emit = defineEmits<{
  (e: 'show-message', message: string, color?: string): void
  (e: 'update-client', client: Client): void
}>()

// Local state
const editMode = ref(false)
const saving = ref(false)
const clientFormRef = ref()
const editableClient = ref<Client | null>(null)   // local copy for editing

const formClient=computed(()=>editableClient.value!)

// Validation rules (same as before)
const rules = {
  required: (v: string) => !!v || 'Required',
  email: (v: string) => !v || /.+@.+\..+/.test(v) || 'Invalid email',
  discount: (v: any) => {
    const num = Number(v)
    return (!isNaN(num) && num >= 0 && num <= 20)
      || 'Discount must be between 0 and 20'
  }
}

// When the selected client changes, reset edit mode and update the local copy
watch(() => props.client, (newClient) => {
  editMode.value = false
  if (newClient) {
    // Clone the client so we don't mutate the prop
    editableClient.value = { ...newClient }
  } else {
    editableClient.value = null
  }
}, { immediate: true })

function toggleEdit() {
  if (editMode.value) {
    // Cancelling: restore original client (keep the prop's value)
    if (props.client) {
      editableClient.value = { ...props.client }
    }
    editMode.value = false
  } else {
    // Enter edit mode: ensure we have a fresh copy of the current client
    if (props.client) {
      editableClient.value = { ...props.client }
    }
    editMode.value = true
  }
}

function cancelEdit() {
  toggleEdit()   // resets form and exits edit mode
}

async function saveClient() {
  // Validate the form
  const { valid } = await clientFormRef.value?.validate()
  if (!valid) return

  // Emit the local copy (which contains the edits)
  if (editableClient.value) {
    saving.value = true
    emit('update-client', editableClient.value)
    saving.value = false
  }
}

// Expose toggleEdit so the parent can manually exit edit mode after a successful save
defineExpose({ toggleEdit })
</script>