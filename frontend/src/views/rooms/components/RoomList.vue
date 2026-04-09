<template>
  <v-toolbar color="#79AE6F" title="Rooms">
    <v-btn 
      icon
      @click="$emit('add')"
      >
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn 
      icon 
      :disabled="!selectedRoomId"
      @click="$emit('delete')"
    >
      <v-icon>mdi-delete</v-icon>
      
    </v-btn>
  </v-toolbar>

  <v-divider />

    <div class="pa-2" style="max-height: 70vh; overflow-y: auto;">
      <template v-for="floor in floors" :key="floor.title">
        <!-- Divider with floor title -->
        <v-divider class="flex-grow-1"></v-divider>
        <h3 class="flex-grow-1 bg-green-lighten-3 my-2 px-3 py-1">{{ floor.title }}</h3>
          <!-- <v-divider class="flex-grow-1"></v-divider> -->
        <v-row dense>
          <v-col
            v-for="room in floor.rooms"
            :key="room.id"
            cols="12"
            sm="6"
            md="4"
            lg="3"
          >
            <v-card
              :variant="selectedRoomId === room.id ? 'tonal' : 'outlined'"
              :color="selectedRoomId === room.id ? 'green' : undefined"
              @click="$emit('select', room)"
              class="cursor-pointer"
            >
              <v-card-title class="text-subtitle-1 font-weight-bold">
                Room {{ room.roomNo }}
              </v-card-title>
              <v-card-text>
                <div class="d-flex justify-space-between align-center mt-2">
                  <span class="text-h6 font-weight-bold text-primary">${{ room.roomCharges }}</span>
                  <v-chip close :color="room.status == 'Available' ? 'green' : 'orange'">
                    {{ room.status || 'Available' }}
                  </v-chip>
                </div>
                <v-divider class="mt-2"></v-divider>
                <div class="d-flex align-center mb-1">                  
                  <v-tooltip v-if="room.doubleBeds" text="Double Beds" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small">mdi-bed-double</v-icon>
                      <span class="ml-1 text-caption">{{ room.doubleBeds }}</span>
                    </template>
                  </v-tooltip>
                  <v-spacer></v-spacer>
                  <v-tooltip v-if="room.singleBeds" text="Single Beds" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small">mdi-bed-single</v-icon>
                      <span class="ml-1 text-caption">{{ room.singleBeds }}</span>
                    </template>
                  </v-tooltip>
                  <v-spacer></v-spacer>
                  <v-tooltip v-if="room.windows" text="Windows" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small">mdi-window-maximize</v-icon>
                      <span class="ml-1 text-caption">{{ room.windows }}</span>
                    </template>
                  </v-tooltip>
                </div>
                <v-divider class="mt-2"></v-divider>
                <div class="d-flex flex-wrap align-center ga-1 mt-1">
                  <v-tooltip v-if="room.aC" text="Air Conditioning" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small"  color="primary">mdi-air-conditioner</v-icon>
                    </template>
                  </v-tooltip>
                  <v-spacer></v-spacer>
                  <v-tooltip v-if="room.wifi" text="Wi-Fi" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small" color="info">mdi-wifi</v-icon>
                    </template>
                  </v-tooltip>
                  <v-spacer></v-spacer>
                  <v-tooltip v-if="room.hotWater" text="Hot Water" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small" color="orange">mdi-water-boiler</v-icon>
                    </template>
                  </v-tooltip>
                  <v-spacer></v-spacer>
                  <v-tooltip v-if="room.balcony" text="Balcony" location="bottom">
                    <template v-slot:activator="{ props }">
                      <v-icon v-bind="props" size="small" color="success">mdi-balcony</v-icon>
                    </template>
                  </v-tooltip>
                </div>
                <!-- Charges and status -->
                <v-divider class="mt-2"></v-divider>
                <div v-if="room.location" class="text-caption mt-1">
                  <v-icon size="x-small">mdi-map-marker</v-icon> {{ room.location }}
                </div>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </template>

      <v-progress-linear v-if="loading" indeterminate />
      <v-alert v-if="error && rooms.length === 0" type="error">{{ error }}</v-alert>
      <v-alert v-if="!loading && rooms.length === 0 && !error" type="info">
        No rooms found
      </v-alert>
    </div>

</template>

<script setup lang="ts">
  import type {Room} from '@/types/index'

  defineProps<{
    rooms:Room[]
    floors:Array<{ title: string; rooms: Room[] }>
    selectedRoomId?: number | null
    loading: boolean
    error?: string | null
  }>()

  defineEmits<{
    (e: 'select', room: Room): void
    (e: 'add'): void
    (e: 'delete'): void
  }>()
  
</script>