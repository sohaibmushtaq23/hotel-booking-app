import { createRouter, createWebHistory } from 'vue-router'
import DashboardView from '@/views/dashboard/DashboardView.vue'
import RoomsView from '@/views/rooms/RoomsView.vue'
import BookingsView from '@/views/bookings/BookingsView.vue'
import ClientsView from '@/views/clients/ClientsView.vue'
import UsersView from '@/views/users/UsersView.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/dashboard' },
    { path: '/dashboard', name: 'dashboard', component: DashboardView },
    { path: '/rooms', name: 'rooms', component: RoomsView },
    { path: '/bookings', name: 'bookings', component: BookingsView },
    { path: '/clients', name: 'clients', component: ClientsView },
    { path: '/users', name: 'users', component: UsersView },
  ],
})

export default router
