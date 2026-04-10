import api from './api';
import type { Booking, BookingDetails } from '@/types';


export const BookingAPI={
  async getBookings(): Promise<BookingDetails[]>{
    const response = await api.get('/reservations');
    return response.data;
  },

  async getBookingById(id: number): Promise<Booking> {
    const response = await api.get(`/reservations/${id}`);
    return response.data;
  },

  async createBooking(bookingData: Omit<Booking, 'id'>): Promise<Booking> {
    const response = await api.post('/reservations', bookingData);
    return response.data;
  },

  async updateBooking(id: number, bookingData: Partial<Booking>): Promise<Booking> {
    const response = await api.put(`/reservations/${id}`, bookingData);
    return response.data;
  },

  async deleteBooking(id: number): Promise<void> {
    await api.delete(`/reservations/${id}`);
  }
}

