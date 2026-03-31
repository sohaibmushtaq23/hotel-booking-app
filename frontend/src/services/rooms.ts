import api from './api';
import type { Room } from '@/types';


export const RoomAPI={
  async getRooms(): Promise<Room[]>{
    const response = await api.get('/rooms');
    return response.data;
  },

  async getRoomById(id: number): Promise<Room> {
    const response = await api.get(`/rooms/${id}`);
    return response.data;
  },

  async createRoom(roomData: Omit<Room, 'id'>): Promise<Room> {
    const response = await api.post('/rooms', roomData);
    return response.data;
  },

  async updateRoom(id: number, roomData: Partial<Room>): Promise<Room> {
    const response = await api.put(`/rooms/${id}`, roomData);
    return response.data;
  },

  async deleteRoom(id: number): Promise<void> {
    await api.delete(`/rooms/${id}`);
  }
}