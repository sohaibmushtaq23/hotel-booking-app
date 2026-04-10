import { defineStore } from 'pinia'
import { ref } from 'vue'
import { BookingAPI } from '@/services/bookings'
import type { Booking, BookingDetails } from '@/types'

export const useBookingStore = defineStore('booking', () => {
  // For the list view – contains joined data (customerName, roomNo, etc.)
  const bookingsList = ref<BookingDetails[]>([])
  // For editing/creating – raw foreign keys
  const selectedBookingRaw = ref<Booking | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchBookings() {
    loading.value = true
    error.value = null
    try {
      const data = await BookingAPI.getBookings() // returns BookingDetails[]
      bookingsList.value = data
      // Auto‑select first if none selected (optional)
      if (data.length > 0 && !selectedBookingRaw.value) {
        // Convert the first details item to a raw Booking if you need to edit it
        const first = data[0]
        if (first){
          selectedBookingRaw.value = {
            id: first.id,
            idCustomer: first.idCustomer,    // you must have this field in BookingDetails
            idRoom: first.idRoom,
            bookingStart: first.bookingStart,
            bookingEnd: first.bookingEnd,
            extraCharges: first.extraCharges,
            amountPaid: first.amountPaid,
            reservedAt: first.reservedAt,
            idReservedBy: first.idReservedBy,
            status: first.status
          }
        }
      }
    } catch (err) {
      error.value = 'Failed to load bookings'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  async function createBooking(bookingData: Omit<Booking, 'id'>) {
    loading.value = true
    error.value = null
    try {
      const newBooking = await BookingAPI.createBooking(bookingData) // returns raw Booking
      // Optionally re‑fetch the list to get the new details (simplest)
      await fetchBookings()
      // Or manually add to bookingsList by fetching details for this new id
      return newBooking
    } catch (err) {
      error.value = 'Failed to create booking'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateBooking(id: number, bookingData: Partial<Booking>) {
    loading.value = true
    error.value = null
    try {
      const updated = await BookingAPI.updateBooking(id, bookingData) // raw Booking
      // Refresh the list to get updated details
      await fetchBookings()
      if (selectedBookingRaw.value?.id === id) {
        selectedBookingRaw.value = updated
      }
      return updated
    } catch (err) {
      error.value = 'Failed to update booking'
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteBooking(id: number) {
    loading.value = true
    error.value = null

    const index = bookingsList.value.findIndex(b => b.id === id)
    if (index === -1) return

    try {
      await BookingAPI.deleteBooking(id)
      bookingsList.value = bookingsList.value.filter(b => b.id !== id)

      if (bookingsList.value.length > 0) {
        let newIndex = index
        if (newIndex >= bookingsList.value.length) {
          newIndex = bookingsList.value.length - 1
        }

        const next = bookingsList.value[newIndex];
        if (next) {
          selectedBookingRaw.value = {
            id: next.id,
            idCustomer: next.idCustomer,
            idRoom: next.idRoom,
            bookingStart: next.bookingStart,
            bookingEnd: next.bookingEnd,
            extraCharges: next.extraCharges,
            amountPaid: next.amountPaid,
            reservedAt: next.reservedAt,
            idReservedBy: next.idReservedBy,
            status: next.status,
          };
        } else {
          selectedBookingRaw.value = null;
        }
      } else {
        selectedBookingRaw.value = null
      }
    } catch (err) {
      error.value = 'Failed to delete booking'
      throw err
    } finally {
      loading.value = false
    }
  }

  function selectBookingForEdit(booking: Booking | null) {
    selectedBookingRaw.value = booking
  }

  return {
    bookingsList,           // use this in the table
    selectedBookingRaw,    // use this in edit/create forms
    loading,
    error,
    fetchBookings,
    createBooking,
    updateBooking,
    deleteBooking,
    selectBookingForEdit
  }
})