<script setup lang="ts">
    import { ref, watch } from 'vue'
    const drawer = ref(false)
    const group = ref(null)

    watch(group, () => {
        drawer.value = false
    })

</script>

<template>
    <v-app>
        <v-app-bar color="#346739">
            <template v-slot:prepend>
                <v-app-bar-nav-icon  variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            </template>
            <v-app-bar-title>Hotel Booking App</v-app-bar-title>
            <v-img
                src="@/assets/logo.png"
                max-height="40"
                max-width="40"
                contain
                class="mr-2"
                >
            </v-img>
        </v-app-bar>
        <v-navigation-drawer color="#346739"
            v-model="drawer"
            temporary
        >
            <v-list>
                <v-list-item prepend-icon="mdi-view-dashboard" title="Dashboard" value="dashboard" to="/"></v-list-item>
                <v-list-item prepend-icon="mdi-book-check" title="Bookings" value="bookings" to="/bookings"></v-list-item>
                <v-list-item prepend-icon="mdi-room-service" title="Rooms" value="rooms" to="/rooms"></v-list-item>
                <v-list-item prepend-icon="mdi-account-group" title="Clients" value="clients" to="/clients"></v-list-item>
                <v-list-item prepend-icon="mdi-account-multiple" title="Users" value="users" to="/users"></v-list-item>
            </v-list>
        </v-navigation-drawer>

    <v-main>
      <v-container fluid >
        <slot/>
      </v-container>
    </v-main>

    <v-footer></v-footer>
  </v-app>
</template>


  