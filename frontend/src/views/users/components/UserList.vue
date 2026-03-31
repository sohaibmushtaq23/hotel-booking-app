<template>
  <v-toolbar color="#79AE6F" title="Users">
    <v-btn 
      icon
      @click="$emit('add')"
      >
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn 
      icon 
      :disabled="!selectedUserId"
      @click="$emit('delete')"
    >
      <v-icon>mdi-delete</v-icon>
      
    </v-btn>
  </v-toolbar>

  <v-divider />

  <!-- Optional header row -->
  <v-list-item class="font-weight-bold text-subtitle-1" style="border-bottom: 1px solid #e0e0e0;">
    <v-row>
      <v-col cols="4">User Name</v-col>
      <v-col cols="4">User Role</v-col>
      <v-col cols="4">Password</v-col>
    </v-row>
  </v-list-item>
  <div class="text-body-2">
    <v-list density="compact">
        <v-list-item
          v-for="user in users"
          :key="user.id"
          :active="selectedUserId === user.id"
          :class="{ 'bg-grey-lighten-2': selectedUserId === user.id }"
          @click="$emit('select', user)"
        >
          <v-row align="center">
            <v-col cols="4">
              <div class="font-weight-bold">{{ user.userName }}</div>
            </v-col>
            <v-col cols="4">{{ user.userRole }}</v-col>
            <v-col cols="4">{{ user.password || '—' }}</v-col>
          </v-row>
        </v-list-item>
      </v-list>
    </div>

</template>

<script setup lang="ts">
  import type {User} from '@/types/index'

  defineProps<{
    users:User[]
    selectedUserId?: number | null
    loading: boolean
    error?: string | null
  }>()

  defineEmits<{
    (e: 'select', user: User): void
    (e: 'add'): void
    (e: 'delete'): void
  }>()
  
</script>