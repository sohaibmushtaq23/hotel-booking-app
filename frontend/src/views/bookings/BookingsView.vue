<template>
    <v-container fluid>
      <v-card>
        <v-toolbar color="#346739" title="Bookings" dark>
          <v-btn color="white" variant="text" prepend-icon="mdi-plus" @click="openCreateDialog">
            New Booking
          </v-btn>
        </v-toolbar>
  
        <!-- Filter chips -->
        <v-card-text>
          <v-chip-group v-model="filterStatus" mandatory>
            <v-chip value="all">All</v-chip>
            <v-chip value="active">Active</v-chip>
            <v-chip value="upcoming">Upcoming</v-chip>
            <v-chip value="past">Past</v-chip>
          </v-chip-group>
        </v-card-text>
  
        <v-divider />
  
        <!-- Data table -->
        <v-data-table
          :items="filteredBookings"
          :loading="store.loading"
          :headers="headers"
          items-per-page="10"
          hover
          @click:row="handleRowClick"
        >
          <template v-slot:item.bookingStart="{ value }">
            {{ value ? new Date(value).toLocaleDateString() : '—' }}
          </template>
          <template v-slot:item.bookingEnd="{ value }">
            {{ value ? new Date(value).toLocaleDateString() : '—' }}
          </template>
          <template v-slot:item.amountPaid="{ value }">
            ${{ value.toFixed(2) }}
          </template>
          <template v-slot:item.status="{ value }">
            <v-chip :color="getStatusColor(value)" size="small">{{ value }}</v-chip>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon size="small" class="me-2" @click.stop="editBooking(item)">
              mdi-pencil
            </v-icon>
            <v-icon size="small" color="error" @click.stop="confirmDelete(item.id)">
              mdi-delete
            </v-icon>
          </template>
        </v-data-table>
      </v-card>
  
      <!-- Create/Edit Dialog -->
      <BookingForm
        v-model="dialogVisible"
        :booking="editingBooking"
        @save="handleSave"
      />
  
      <!-- Delete confirmation dialog -->
      <v-dialog v-model="deleteDialog.show" max-width="400">
        <v-card title="Confirm Delete">
          <v-card-text>Delete booking #{{ deleteDialog.bookingId }}?</v-card-text>
          <v-card-actions>
            <v-spacer />
            <v-btn @click="deleteDialog.show = false">Cancel</v-btn>
            <v-btn color="error" @click="confirmDeleteBooking">Delete</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
  
      <!-- Snackbar -->
      <v-snackbar v-model="snackbar.show" :color="snackbar.color" timeout="3000">
        {{ snackbar.message }}
      </v-snackbar>
    </v-container>
  </template>
  
  <script setup lang="ts">
  import { ref, computed, onMounted } from 'vue'
  import { useBookingStore } from '@/stores/bookingStore'
  import type { Booking, BookingDetails } from '@/types'
  import BookingForm from './components/BookingForm.vue'
  
  const store = useBookingStore()
  const filterStatus = ref('all') // 'all', 'active', 'upcoming', 'past'
  
  // Table headers
  const headers = [
    { title: 'Customer', key: 'customerName' },
    { title: 'Room', key: 'roomNo' },
    { title: 'Check-in', key: 'bookingStart' },
    { title: 'Check-out', key: 'bookingEnd' },
    { title: 'Paid', key: 'amountPaid', align: 'end' },
    { title: 'Status', key: 'status' },
    { title: 'Actions', key: 'actions', sortable: false },
  ] as const
  
  // Filter logic
  const today = new Date().toISOString().slice(0, 10)
  
  const filteredBookings = computed(() => {
    const list = store.bookingsList
    if (filterStatus.value === 'all') return list
  
    return list.filter(booking => {
      const start = booking.bookingStart?.slice(0, 10) || ''
      const end = booking.bookingEnd?.slice(0, 10) || ''
      const status = booking.status
  
      switch (filterStatus.value) {
        case 'active':
          return status === 'Confirmed' && start <= today && end >= today
        case 'upcoming':
          return status === 'Confirmed' && start > today
        case 'past':
          return status === 'Completed' || (end && end < today)
        default:
          return true
      }
    })
  })
  
  function getStatusColor(status: string): string {
    switch (status) {
      case 'Confirmed': return 'success'
      case 'Pending': return 'warning'
      case 'Completed': return 'grey'
      case 'Cancelled': return 'error'
      default: return 'info'
    }
  }
  
  // Dialog handling
  const dialogVisible = ref(false)
  const editingBooking = ref<Booking | null>(null)
  
  function openCreateDialog() {
    editingBooking.value = null
    dialogVisible.value = true
  }
  
  function editBooking(booking: BookingDetails) {
    // Convert BookingDetails to raw Booking for editing
    editingBooking.value = {
      id: booking.id,
      idCustomer: booking.idCustomer,
      idRoom: booking.idRoom,
      bookingStart: booking.bookingStart,
      bookingEnd: booking.bookingEnd,
      extraCharges: booking.extraCharges,
      amountPaid: booking.amountPaid,
      reservedAt: booking.reservedAt,
      idReservedBy: booking.idReservedBy,
      status: booking.status
    }
    dialogVisible.value = true
  }
  
  async function handleSave(bookingData: Omit<Booking, 'id'> | Booking) {
    try {
      if (editingBooking.value?.id) {
        await store.updateBooking(editingBooking.value.id, bookingData)
        showNotification('Booking updated')
      } else {
        await store.createBooking(bookingData as Omit<Booking, 'id'>)
        showNotification('Booking created')
      }
      dialogVisible.value = false
    } catch (err) {
      showNotification('Operation failed', 'error')
    }
  }
  
  // Delete handling
  const deleteDialog = ref({ show: false, bookingId: 0 })
  function confirmDelete(id: number) {
    deleteDialog.value = { show: true, bookingId: id }
  }
  async function confirmDeleteBooking() {
    try {
      await store.deleteBooking(deleteDialog.value.bookingId)
      showNotification('Booking deleted')
    } catch {
      showNotification('Delete failed', 'error')
    } finally {
      deleteDialog.value.show = false
    }
  }
  
  // Snackbar
  const snackbar = ref({ show: false, message: '', color: 'success' })
  function showNotification(msg: string, color = 'success') {
    snackbar.value = { show: true, message: msg, color }
  }
  
  // Row click: show details in a separate panel (optional)
  function handleRowClick(event: any, { item }: { item: BookingDetails }) {
    // You could open a side panel or dialog with full details
    console.log('Selected booking', item)
  }
  
  onMounted(() => {
    store.fetchBookings()
  })
  </script>