<template>
    <v-card>
      <v-toolbar flat color="#79AE6F">
        <v-toolbar-title>Clients</v-toolbar-title>
        <v-spacer />
        <v-btn icon @click="$emit('add')">
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
  
      <v-list v-if="!loading" density="compact">
        <v-list-item
          v-for="client in clients"
          :key="client.id"
          :title="client.clientName"
          :subtitle="'CNIC:'+client.cnic+ ' Email:' +client.email"
          :active="selectedClientId === client.id"
          @click="$emit('select', client)"
        />
      </v-list>
      <v-progress-linear v-if="loading" indeterminate />
      <v-alert v-if="error" type="error">{{ error }}</v-alert>
    </v-card>
  </template>
  
  <script setup lang="ts">
  import type { Client } from '@/types'
  
  defineProps<{
    clients: Client[]
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