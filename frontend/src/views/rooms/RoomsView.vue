<template>
    <v-container fluid fill-height>
      <v-row>
        <!-- Left column: rooms list -->
        <v-col cols="6">
          <RoomList
            :rooms="store.rooms"
            :floors="floors"
            :selected-room-id="store.selectedRoom?.id"
            :loading="store.loading"
            :error="store.error"
            @select="store.selectRoom"
            @add="openAddDialog"
            @delete="openDeleteDialog(store.selectedRoom!)"
            @create-booking="openBookingDialog"
            />

        </v-col>
  
        <!-- Right column: room details -->
        <v-col cols="6">
          <RoomDetails
            ref="roomDetailsRef"
            :locations="locations"
            :stati="stati"
            :room="store.selectedRoom"
            @show-message="showNotification"
            @update-room="handleUpdateRoom"
          />
        </v-col>
      </v-row>
  
      <!-- Room Delete Dialog -->
      <v-dialog v-model="deleteDialog.show" max-width="400">
        <v-card title="Confirm Delete">
          <v-card-text>
            Are you sure you want to delete {{ deleteDialog.roomNo }}?
          </v-card-text>

          <v-card-actions>
            <v-spacer/>
            <v-btn color="grey" @click="deleteDialog.show=false">Cancel</v-btn>
            <v-btn color="success" @click="confirmDelete" :loading="store.loading">Delete</v-btn>
          </v-card-actions>
          
        </v-card>
      </v-dialog>
  
      <!-- Add Room Dialog -->
      <v-dialog v-model="addDialog" max-width="600">
        <v-card title="Add New Room">
          <v-card-text>
            <v-form ref="addFormRef" v-model="addValid" >
              <v-text-field v-model="newRoom.roomNo" label="Room Number" :rules="[rules.required]"></v-text-field>
              <v-text-field v-model.number="newRoom.roomWidth" label="Room Width (feet)" type="number" step="0.25" min="0" max="50" :rules="[rules.positiveNumber]"></v-text-field>
              <v-text-field v-model.number="newRoom.roomLength" label="Room Length (feet)" type="number" step="0.25" min="0" max="50" :rules="[rules.positiveNumber]"></v-text-field>
              <v-text-field v-model.number="newRoom.doubleBeds" label="Double Beds" type="number" min="0" max="10" :rules="[rules.positiveNumber]"></v-text-field>
              <v-text-field v-model.number="newRoom.singleBeds" label="Single Beds" type="number" min="0" max="10" :rules="[rules.positiveNumber]"></v-text-field>
              <v-text-field v-model.number="newRoom.windows" label="Windows" type="number" min="0" max="10" :rules="[rules.positiveNumber]"></v-text-field>
              <v-checkbox label="AC" v-model="newRoom.aC"></v-checkbox>
              <v-checkbox label="Wifi" v-model="newRoom.wifi"></v-checkbox>
              <v-checkbox label="Hot Water" v-model="newRoom.hotWater"></v-checkbox>
              <v-checkbox label="Balcony" v-model="newRoom.balcony"></v-checkbox>
              <v-select :items="locations" v-model="newRoom.location" label="Location"></v-select>
              <v-text-field v-model.number="newRoom.roomCharges" label="Charges" type="number" :rules="[rules.positiveNumber]"></v-text-field>
              <v-text-field v-model="newRoom.remarks" label="Remarks"></v-text-field>
              <v-select :items="stati" v-model="newRoom.status" label="Status"></v-select>
              
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn color="grey" @click="addDialog=false">Cancel</v-btn>
            <v-btn color="success" @click="saveNewRoom" :loading="store.loading">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Booking Creation Dialog -->
      <BookingForm
        v-model="bookingDialogVisible"
        :booking="editingBookingForDialog"
        @save="onBookingSaved"
      />
  
      <!-- Global Snackbar -->
      <v-snackbar v-model="snackbar.show" :color="snackbar.color" timeout="3000">
        {{ snackbar.message }}
        <template v-slot:actions>
          <v-btn color="white" variant="text" @click="snackbar.show = false">Close</v-btn>
        </template>
      </v-snackbar>
    </v-container>
  </template>
  
  <script setup lang="ts">
    import { ref, computed } from 'vue'
    import { useRoomStore } from '@/stores/roomStore'
    import { useBookingStore } from '@/stores/bookingStore'
    import type {Booking, Room } from '@/types'
    import RoomList from './components/RoomList.vue'
    import RoomDetails from './components/RoomDetails.vue'
    import BookingForm from '@/views/bookings/components/BookingForm.vue'

    const store=useRoomStore()
    const bookingStore=useBookingStore()
    store.fetchRooms()

    const floorNames = ['Ground', 'First', 'Second', 'Third']

    // Group rooms by floor (using location string)
    const floors = computed(() => {
      // Create a map: floor name -> array of rooms
      const map = new Map<string, Room[]>()

      // Initialize empty arrays for each floor
      floorNames.forEach(floor => map.set(floor, []))

      // Group rooms
      store.rooms.forEach(room => {
        const loc = room.location?.toLowerCase() || ''
        let floorKey: string | null = null

        if (loc.includes('ground')) floorKey = 'Ground'
        else if (loc.includes('first')) floorKey = 'First'
        else if (loc.includes('second')) floorKey = 'Second'
        else if (loc.includes('third')) floorKey = 'Third'
        else floorKey = null // skip unknown locations or assign to an "Other" group

        if (floorKey && map.has(floorKey)) {
          map.get(floorKey)!.push(room)
        }
        // If you want an "Other" floor, add it to floorNames and handle here
      })

      // Convert map to array of objects for easier template iteration
      return floorNames
        .filter(floor => map.get(floor)!.length > 0) // optional: skip empty floors
        .map(floor => ({
          title: `${floor} Floor`,
          rooms: map.get(floor)!
        }))
    })

    //===Add dialog===
    const addDialog=ref(false)
    const addFormRef=ref()
    const addValid=ref(false)
    const roomDetailsRef = ref()

    const newRoom=ref<Omit<Room,'id'>>(
      {
        roomNo:'',
        roomWidth:0,
        roomLength:0,
        doubleBeds:0,
        singleBeds:0,
        windows:0,
        aC:false,
        wifi:false,
        hotWater: false,
        balcony:false,
        location:'',
        roomCharges: 0,
        roomImage:'',
        remarks:'',
        status:''
      }
    )

    const locations=["Ground Floor","First Floor", "Second Floor","Third Floor"]
    const stati=["Available","Reserved"]

    //Validation Rules
    const rules = {
      required: (v: string) => !!v || 'Required',
      email: (v: string) => !v || /.+@.+\..+/.test(v) || 'Invalid email',
      positiveNumber: (v: any) => {
        const num = Number(v)
        return (!isNaN(num) && num >= 0)
          || 'Value must be non-negative'
      }
    }

    function openAddDialog(){
      newRoom.value={
        roomNo:'',
        roomWidth:0,
        roomLength:0,
        doubleBeds:0,
        singleBeds:0,
        windows:0,
        aC:false,
        wifi:false,
        hotWater: false,
        balcony:false,
        location:'',
        roomCharges: 0,
        roomImage:'',
        remarks:'',
        status:''
      }
      addDialog.value=true
    }

    async function saveNewRoom(){
      const { valid } = await addFormRef.value.validate()
      if (!valid) return

      try{
        await store.createRoom(newRoom.value)
        addDialog.value=false
        showNotification('Room created successfully')

      }catch(err){
        showNotification('Failed to create the room','error')

      }
    }

    //===Delete dialog===
    const deleteDialog=ref({
      show: false,
      roomId:0,
      roomNo:''
    })

    function openDeleteDialog(room: Room){

      deleteDialog.value={
        show: true,
        roomId:room.id,
        roomNo:room.roomNo
      }
    }

    async function confirmDelete() {
      try{
        await store.deleteRoom(deleteDialog.value.roomId)
        deleteDialog.value.show=false
        showNotification('Room deleted successfully')

      }catch(err){
        showNotification('Failed to delete the room','error')

      }
      
    }

    //===Update Room===
    async function handleUpdateRoom(updatedRoom: Room) {
      try {
        await store.updateRoom(updatedRoom.id, updatedRoom)
        showNotification('Room updated successfully')
      } catch (err) {
        showNotification('Failed to update room', 'error')
      }
    }

    //===Snackbar===
    const snackbar=ref({
      show: false,
      message:'',
      color:'success'
    })

    function showNotification(message: string, color= 'success'){
      snackbar.value={show:true, message, color}
    }

    const bookingDialogVisible = ref(false)
    const editingBookingForDialog = ref<Booking | Omit<Booking, 'id'> | null>(null)

    function openBookingDialog(room: Room) {
      // Create a new booking object (no id) with the room preselected
      editingBookingForDialog.value = {
        idCustomer: 0,
        idRoom: room.id,
        bookingStart: '',
        bookingEnd: '',
        extraCharges: 0,
        amountPaid: 0,
        reservedAt: null,
        idReservedBy: 1004,
        status: 'Pending'
      }
      bookingDialogVisible.value = true
    }

    async function onBookingSaved(bookingData: Booking | Omit<Booking, 'id'>) {
      try {
        if ('id' in bookingData && bookingData.id) {
          // Update existing booking
          await bookingStore.updateBooking(bookingData.id, bookingData)
          showNotification('Booking updated successfully')
        } else {
          // Create new booking
          await bookingStore.createBooking(bookingData)
          showNotification('Booking created successfully')
        }
        // Refresh room list to reflect updated status (e.g., from Available to Reserved)
        await store.fetchRooms()
      } catch (err) {
        showNotification('Operation failed', 'error')
      } finally {
        bookingDialogVisible.value = false
        editingBookingForDialog.value = null
      }
    }

  </script>