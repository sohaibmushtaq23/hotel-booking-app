<template>
    <v-card>
      <v-toolbar flat color="#79AE6F" dark dense>
        <v-toolbar-title>Client Details</v-toolbar-title>
        <v-spacer />
  
        <!-- Edit/Cancel and Save buttons on Details tab -->
        <template v-if="activeTab === 0 && client">
          <!-- Edit / Cancel button -->
          <v-tooltip text="Edit" location="bottom">
            <template v-slot:activator="{ props }">
              <v-btn
                v-if="!editMode"
                v-bind="props"
                icon
                color="white"
                @click="toggleEdit"
              >
                <v-icon>mdi-pencil</v-icon>
              </v-btn>
            </template>
          </v-tooltip>
  
          <v-tooltip text="Cancel" location="bottom">
            <template v-slot:activator="{ props }">
              <v-btn
                v-if="editMode"
                v-bind="props"
                icon
                color="white"
                @click="toggleEdit"
              >
                <v-icon>mdi-close</v-icon>
              </v-btn>
            </template>
          </v-tooltip>
  
          <!-- Save button (only in edit mode) -->
          <v-tooltip text="Save" location="bottom">
            <template v-slot:activator="{ props }">
              <v-btn
                v-if="editMode"
                v-bind="props"
                icon
                color="white"
                class="ml-2"
                @click="saveClient"
                :loading="clientStore.loading"
              >
                <v-icon>mdi-check</v-icon>
              </v-btn>
            </template>
          </v-tooltip>
        </template>
  
        <!-- Add Contact button on Contacts tab -->
        <template v-if="activeTab === 1 && client">
          <v-tooltip text="Add Contact" location="bottom">
            <template v-slot:activator="{ props }">
              <v-btn
                v-bind="props"
                icon
                color="white"
                @click="contactsManagerRef?.openAddDialog()"
              >
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </template>
          </v-tooltip>
        </template>
      </v-toolbar>
  
      <v-tabs v-model="activeTab" background-color="primary" dark>
        <v-tab>Details</v-tab>
        <v-tab>Contacts</v-tab>
      </v-tabs>
  
      <v-tabs-window v-model="activeTab">
        <!-- Details Tab -->
        <v-tabs-window-item>
          <v-card-text>
            <div v-if="!client" class="text-center text-grey">
              Select a client to view details
            </div>
            <ClientForm
              v-else
              ref="clientFormRef"
              :client="client"
              :edit-mode="editMode"
            />
          </v-card-text>
        </v-tabs-window-item>
  
        <!-- Contacts Tab -->
        <v-tabs-window-item>
          <v-card-text>
            <div v-if="!client" class="text-center text-grey">
              Select a client to view contacts
            </div>
            <ContactsManager
              v-else
              ref="contactsManagerRef"
              :client-id="client.id"
              @show-message="emit('show-message', $event)"
            />
          </v-card-text>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card>
  </template>
  
  <script setup lang="ts">
  import { ref, watch } from 'vue'
  import { useClientStore } from '@/stores/clientStore'
  import type { Client } from '@/types'
  import ClientForm from './ClientForm.vue'
  
  const props = defineProps<{
    client: Client | null
  }>()
  
  const emit = defineEmits<{
    (e: 'show-message', message: string, color?: string): void
  }>()
  
  const clientStore = useClientStore()
  
  // Local state
  const editMode = ref(false)
  const activeTab = ref(0) // 0 = Details, 1 = Contacts
  const clientFormRef = ref()
  const contactsManagerRef = ref()
  
  // Watch for client changes to reset edit mode
  watch(() => props.client, () => {
    editMode.value = false
  }, { immediate: true })
  
  function toggleEdit() {
    if (editMode.value) {
      // Cancel: reset form and exit edit mode
      clientFormRef.value?.reset()
      editMode.value = false
    } else {
      // Enter edit mode
      editMode.value = true
    }
  }
  
  async function saveClient() {
    const isValid = await clientFormRef.value?.validate()
    if (!isValid) return
  
    const updatedClient = clientFormRef.value?.getData()
    try {
      await clientStore.updateClient(updatedClient.id, updatedClient)
      editMode.value = false
      emit('show-message', 'Client updated successfully')
    } catch (err) {
      emit('show-message', 'Failed to update client', 'error')
    }
  }
  </script>