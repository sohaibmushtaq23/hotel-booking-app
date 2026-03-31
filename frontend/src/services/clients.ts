import api from './api';
import type { Client } from '@/types';

export const ClientAPI = {
  async getClients(): Promise<Client[]> {
    const response = await api.get('/clients');
    return response.data;
  },

  async getClientById(id: number): Promise<Client> {
    const response = await api.get(`/clients/${id}`);
    return response.data;
  },

  async createClient(clientData: Omit<Client, 'id'>): Promise<Client> {
    const response = await api.post('/clients', clientData);
    return response.data;
  },

  async updateClient(id: number, clientData: Partial<Client>): Promise<Client> {
    const response = await api.put(`/clients/${id}`, clientData);
    return response.data;
  },

  async deleteClient(id: number): Promise<void> {
    await api.delete(`/clients/${id}`);
  }
};