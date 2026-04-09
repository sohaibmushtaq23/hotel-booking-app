import { defineStore } from 'pinia'
import { ref } from 'vue'
import { RoomAPI } from '@/services/rooms'
import type { Room } from '@/types'

export const useRoomStore = defineStore('room', () => {
  const rooms = ref<Room[]>([])
  const selectedRoom = ref<Room | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchRooms() {
    loading.value = true
    error.value = null
    try {
      const data = await RoomAPI.getRooms()
      rooms.value = data

    // Auto‑select first room if list not empty and none selected
    if (data.length > 0 && !selectedRoom.value) {
      selectRoom(data[0]!)
    }
    } catch (err) {
      error.value = 'Failed to load rooms'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  async function createRoom(roomData: Omit<Room, 'id'>) {
    loading.value = true
    error.value = null
    try {
      const newRoom = await RoomAPI.createRoom(roomData)
      rooms.value.push(newRoom)

      selectRoom(newRoom)
      return newRoom
    } catch (err) {
      error.value = 'Failed to create room'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateRoom(id: number, roomData: Partial<Room>) {
    loading.value = true
    error.value = null
    try {
      const updated = await RoomAPI.updateRoom(id, roomData)
      console.log('Updating room, response:', updated)
      // Update in local list
      const index = rooms.value.findIndex(c => c.id === id)
      if (index !== -1) {
        rooms.value[index] = updated
      }

      if (selectedRoom.value?.id === id) {
        selectedRoom.value = updated
      }
      return updated
    } catch (err) {
      error.value = 'Failed to update room'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteRoom(id: number) {
    loading.value = true
    error.value = null
  
    // Find the index of the room to delete (before removal)
    const index = rooms.value.findIndex(c => c.id === id)
    if (index === -1) return  // not found, nothing to delete
  
    try {
      await RoomAPI.deleteRoom(id)
  
      // Remove from local list
      rooms.value = rooms.value.filter(c => c.id !== id)
  
      if (rooms.value.length > 0) {
        let newIndex = index
        if (newIndex >= rooms.value.length) {
          newIndex = rooms.value.length - 1
        }
        selectRoom(rooms.value[newIndex]!)
      } else {

        selectedRoom.value = null
      }
    } catch (err) {
      error.value = 'Failed to delete room'
      throw err
    } finally {
      loading.value = false
    }
  }

  function selectRoom(room: Room | null) {
    selectedRoom.value = room
  }

  return {
    rooms,
    selectedRoom,
    loading,
    error,
    fetchRooms,
    createRoom,
    updateRoom,
    deleteRoom,
    selectRoom
  }
})