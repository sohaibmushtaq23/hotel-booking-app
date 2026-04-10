<template>
    <v-dialog v-model="visible" max-width="800" persistent>
      <v-card>
        <v-card-title>{{ isEdit ? 'Edit Booking' : 'New Booking' }}</v-card-title>
        <v-card-text>
          <v-form ref="formRef" v-model="valid">
            <v-row>
              <v-col cols="6">
                <v-select
                  v-model="formData.idCustomer"
                  :items="clients"
                  item-title="clientName"
                  item-value="id"
                  label="Customer"
                  :rules="[rules.required]"
                  :loading="clientStore.loading"
                />
              </v-col>
              <v-col cols="6">
                <v-select
                  v-model="formData.idRoom"
                  :items="rooms"
                  item-title="roomNo"
                  item-value="id"
                  label="Room"
                  :rules="[rules.required]"
                  :loading="roomStore.loading"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-text-field
                  v-model="formData.bookingStart"
                  label="Check-in"
                  type="datetime-local"
                  :rules="[rules.required]"
                />

              </v-col>
              <v-col cols="6">
                <v-text-field
                  v-model="formData.bookingEnd"
                  label="Check-out"
                  type="datetime-local"
                  :rules="[rules.required, rules.endAfterStart]"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-text-field
                  v-model.number="formData.amountPaid"
                  label="Amount Paid"
                  type="number"
                  prefix="$"
                  :rules="[rules.positive]"
                />
              </v-col>
              <v-col cols="6">
                <v-text-field
                  v-model.number="formData.extraCharges"
                  label="Extra Charges"
                  type="number"
                  prefix="$"
                  :rules="[rules.positive]"
                />
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-select
                  v-model="formData.status"
                  :items="statusOptions"
                  label="Status"
                  :rules="[rules.required]"
                />
              </v-col>
            </v-row>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn @click="close">Cancel</v-btn>
          <v-btn color="primary" @click="save" :loading="saving">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </template>
  
  <script setup lang="ts">
  import { ref, computed, watch } from 'vue'
  import type { Booking } from '@/types'
  import { useClientStore } from '@/stores/clientStore'
  import { useRoomStore } from '@/stores/roomStore'
  
  const props = defineProps<{
    modelValue: boolean
    booking: Booking | null
  }>()
  
  const emit = defineEmits<{
    (e: 'update:modelValue', value: boolean): void
    (e: 'save', data: Omit<Booking, 'id'> | Booking): void
  }>()
  
  const visible = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val)
  })
  
  const isEdit = computed(() => !!props.booking?.id)
  
  const clientStore = useClientStore()
  const roomStore = useRoomStore()
  
  // Fetch data for dropdowns
  clientStore.fetchClients()
  roomStore.fetchRooms()
  
  const clients = computed(() => clientStore.clients)
  const rooms = computed(() => roomStore.rooms)
  
  const statusOptions = ['Pending', 'Confirmed', 'CheckedIn', 'Completed', 'Cancelled']
  
  const formRef = ref()
  const valid = ref(false)
  const saving = ref(false)
  
  // Form data
  const formData = ref<Omit<Booking, 'id'> & { id?: number }>({
    idCustomer: 0,
    idRoom: 0,
    bookingStart: '',
    bookingEnd: '',
    extraCharges: 0,
    amountPaid: 0,
    idReservedBy: 1,
    reservedAt:null,
    status: 'Pending'
  })
  
  // Rules
  const rules = {
    required: (v: any) => !!v || 'Required',
    positive: (v: number) => v >= 0 || 'Must be ≥ 0',
    endAfterStart: (v: string) => {
      if (!formData.value.bookingStart || !v) return true
      return new Date(v) >= new Date(formData.value.bookingStart) || 'Check-out must be after check-in'
    }
  }
  
  watch(() => props.booking, (newVal) => {
    if (newVal) {
      formData.value = { ...newVal }
    } else {
      formData.value = {
        idCustomer: 0,
        idRoom: 0,
        bookingStart: '',
        bookingEnd: '',
        extraCharges: 0,
        amountPaid: 0,
        idReservedBy: 1004,
        reservedAt:null,
        status: 'Pending'
      }
    }
  }, { immediate: true })
  
  async function save() {
    const { valid: isValid } = await formRef.value.validate()
    if (!isValid) return
    saving.value = true
    try {
      const payload = { ...formData.value,
    bookingStart: toUTCWithoutShift(formData.value.bookingStart),
    bookingEnd: toUTCWithoutShift(formData.value.bookingEnd),
    reservedAt: null,
   }
      if (!isEdit.value) delete payload.id
      emit('save', payload)
      close()
    } finally {
      saving.value = false
    }
  }
  
  function close() {
    visible.value = false
  }

function toUTCWithoutShift(localDateTime: string|null): string | null {
  if (!localDateTime) return null

  return localDateTime + ':00Z'
}

  </script>