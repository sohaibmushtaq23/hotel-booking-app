<template>
    <v-form ref="formRef" v-model="valid">
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="localClient.clientName"
            label="Client Code"
            :readonly="!editMode"
            :rules="[rules.required]"
          />
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="localClient.cnic"
            label="CNIC"
            :readonly="!editMode"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model="localClient.phone"
            label="Phone"
            :readonly="!editMode"
          />
        </v-col>

        <v-col cols="12" md="6">
          <v-text-field
            v-model="localClient.email"
            label="Email"
            :readonly="!editMode"
            :rules="[rules.email]"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
            v-model.number="localClient.discount"
            label="Discount"
            type="number"
            :readonly="!editMode"
          />
        </v-col>
      </v-row>
    </v-form>
  </template>
  
  <script setup lang="ts">
  import { ref, watch } from 'vue'
  import type { Client } from '@/types'
  
  const props = defineProps<{
    client: Client | null
    editMode: boolean
  }>()
  
  // Emit events (we'll use exposed methods instead)
  // defineEmits([])

  const rules = {
    required: (v: string) => !!v || 'Required',
    email: (v: string) => !v || /.+@.+\..+/.test(v) || 'Invalid email'
  }
  
  const localClient = ref<Client>({ ...props.client } as Client)
  const valid = ref(false)
  const formRef = ref()
  
  // Watch for external client changes (when selection changes)
  watch(() => props.client, (newClient) => {
    if (newClient) {
      localClient.value = { ...newClient }
    }
  }, { immediate: true })
  
  // Expose methods to parent
  defineExpose({
    validate: async () => {
      const { valid } = await formRef.value?.validate() ?? { valid: false }
      return valid
    },
    getData: () => localClient.value,
    reset: () => {
      if (props.client) {
        localClient.value = { ...props.client }
      }
    }
  })
  </script>