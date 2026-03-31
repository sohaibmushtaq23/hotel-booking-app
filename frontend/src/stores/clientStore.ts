import { defineStore } from 'pinia'
import { ref } from 'vue'
import { ClientAPI } from '@/services/clients'
import type { Client } from '@/types'

export const useClientStore = defineStore('client', () => {
  const clients = ref<Client[]>([])
  const selectedClient = ref<Client | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchClients() {
    loading.value = true
    error.value = null
    try {
      const data = await ClientAPI.getClients()
      clients.value = data

    // Auto‑select first client if list not empty and none selected
    if (data.length > 0 && !selectedClient.value) {
      selectClient(data[0]!)
    }
    } catch (err) {
      error.value = 'Failed to load clients'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  async function createClient(clientData: Omit<Client, 'id'>) {
    loading.value = true
    error.value = null
    try {
      const newClient = await ClientAPI.createClient(clientData)
      // Option 1: add to local list (optimistic)
      clients.value.push(newClient)
      // Option 2: refetch all (simpler, ensures consistency)
      // await fetchClients()
      selectClient(newClient)
      return newClient
    } catch (err) {
      error.value = 'Failed to create client'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateClient(id: number, clientData: Partial<Client>) {
    loading.value = true
    error.value = null
    try {
      const updated = await ClientAPI.updateClient(id, clientData)
      console.log('Updating client, response:', updated)
      // Update in local list
      const index = clients.value.findIndex(c => c.id === id)
      if (index !== -1) {
        clients.value[index] = updated
      }
      // If this is the selected client, update selectedClient
      if (selectedClient.value?.id === id) {
        selectedClient.value = updated
      }
      return updated
    } catch (err) {
      error.value = 'Failed to update client'
      throw err
    } finally {
      loading.value = false
    }
  }

  // async function deleteClient(id: number) {
  //   loading.value = true
  //   error.value = null
  //   try {
  //     await ClientAPI.deleteClient(id)
  //     // Remove from local list
  //     clients.value = clients.value.filter(c => c.id !== id)
  //     // If deleted client was selected, clear selection
  //     if (selectedClient.value?.id === id) {
  //       selectedClient.value = null
  //     }
  //   } catch (err) {
  //     error.value = 'Failed to delete client'
  //     throw err
  //   } finally {
  //     loading.value = false
  //   }
  // }

  async function deleteClient(id: number) {
    loading.value = true
    error.value = null
  
    // Find the index of the client to delete (before removal)
    const index = clients.value.findIndex(c => c.id === id)
    if (index === -1) return  // not found, nothing to delete
  
    try {
      await ClientAPI.deleteClient(id)
  
      // Remove from local list
      clients.value = clients.value.filter(c => c.id !== id)
  
      // Determine which client to select next (if any)
      if (clients.value.length > 0) {
        let newIndex = index
        // If the removed client was not the last one, the client at the same index becomes the next one
        if (newIndex >= clients.value.length) {
          // Deleted was the last → select the new last client
          newIndex = clients.value.length - 1
        }
        selectClient(clients.value[newIndex]!)
      } else {
        // No clients left → clear selection
        selectedClient.value = null
      }
    } catch (err) {
      error.value = 'Failed to delete client'
      throw err
    } finally {
      loading.value = false
    }
  }

  function selectClient(client: Client | null) {
    selectedClient.value = client
  }

  return {
    clients,
    selectedClient,
    loading,
    error,
    fetchClients,
    createClient,
    updateClient,
    deleteClient,
    selectClient
  }
})