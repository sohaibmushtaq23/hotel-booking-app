import { defineStore } from 'pinia'
import { ref } from 'vue'
import { UserAPI } from '@/services/users'
import type { User } from '@/types'

export const useUserStore = defineStore('user', () => {
  const users = ref<User[]>([])
  const selectedUser = ref<User | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchUsers() {
    loading.value = true
    error.value = null
    try {
      const data = await UserAPI.getUsers()
      users.value = data

    // Auto‑select first user if list not empty and none selected
    if (data.length > 0 && !selectedUser.value) {
      selectUser(data[0]!)
    }
    } catch (err) {
      error.value = 'Failed to load users'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  async function createUser(userData: Omit<User, 'id'>) {
    loading.value = true
    error.value = null
    try {
      const newUser = await UserAPI.createUser(userData)
      users.value.push(newUser)

      selectUser(newUser)
      return newUser
    } catch (err) {
      error.value = 'Failed to create user'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateUser(id: number, userData: Partial<User>) {
    loading.value = true
    error.value = null
    try {
      const updated = await UserAPI.updateUser(id, userData)
      console.log('Updating user, response:', updated)
      // Update in local list
      const index = users.value.findIndex(c => c.id === id)
      if (index !== -1) {
        users.value[index] = updated
      }

      if (selectedUser.value?.id === id) {
        selectedUser.value = updated
      }
      return updated
    } catch (err) {
      error.value = 'Failed to update user'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteUser(id: number) {
    loading.value = true
    error.value = null
  
    // Find the index of the user to delete (before removal)
    const index = users.value.findIndex(c => c.id === id)
    if (index === -1) return  // not found, nothing to delete
  
    try {
      await UserAPI.deleteUser(id)
  
      // Remove from local list
      users.value = users.value.filter(c => c.id !== id)
  
      if (users.value.length > 0) {
        let newIndex = index
        if (newIndex >= users.value.length) {
          newIndex = users.value.length - 1
        }
        selectUser(users.value[newIndex]!)
      } else {

        selectedUser.value = null
      }
    } catch (err) {
      error.value = 'Failed to delete user'
      throw err
    } finally {
      loading.value = false
    }
  }

  function selectUser(user: User | null) {
    selectedUser.value = user
  }

  return {
    users,
    selectedUser,
    loading,
    error,
    fetchUsers,
    createUser,
    updateUser,
    deleteUser,
    selectUser
  }
})