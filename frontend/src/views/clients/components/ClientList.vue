<template>
  <v-toolbar color="#79AE6F" title="Clients">
    <v-btn 
      icon
      @click="$emit('add')"
      >
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn 
      icon 
      :disabled="!selectedClientId"
      @click="$emit('delete')"
    >
      <v-icon>mdi-delete</v-icon>
      
    </v-btn>
  </v-toolbar>

  <v-divider />

  <!-- Optional header row -->
  <v-list-item class="font-weight-bold text-subtitle-1" style="border-bottom: 1px solid #e0e0e0;">
    <v-row>
      <v-col cols="3">Name</v-col>
      <v-col cols="2">CNIC</v-col>
      <v-col cols="2">Phone</v-col>
      <v-col cols="3">Email</v-col>
      <v-col cols="2" class="text-right">Discount</v-col>
    </v-row>
  </v-list-item>
  <div class="text-body-2">
    <v-list density="compact">
        <v-list-item
          v-for="client in clients"
          :key="client.id"
          :active="selectedClientId === client.id"
          :class="{ 'bg-grey-lighten-2': selectedClientId === client.id }"
          @click="$emit('select', client)"
        >
          <v-row align="center">
            <v-col cols="3">
              <div class="font-weight-bold">{{ client.clientName }}</div>
            </v-col>
            <v-col cols="2">{{ client.cnic }}</v-col>
            <v-col cols="2">{{ client.phone || '—' }}</v-col>
            <v-col cols="3">{{ client.email || '—' }}</v-col>
            <v-col cols="2" class="text-right">{{ client.discount }}%</v-col>
          </v-row>
        </v-list-item>
      </v-list>
    </div>

</template>

<script setup lang="ts">
  import type {Client} from '@/types/index'

  defineProps<{
    clients:Client[]
    selectedClientId?: number | null
    loading: boolean
    error?: string | null
  }>()

  defineEmits<{
    (e: 'select', client: Client): void
    (e: 'add'): void
    (e: 'delete'): void
  }>()
  
</script>