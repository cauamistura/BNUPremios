import { authenticatedFetch } from './apiUtils';

export const usersService = {
  async listUsers() {
    const res = await authenticatedFetch('/users', { method: 'GET' });
    return res.json();
  },
  async getUserById(id: string) {
    const res = await authenticatedFetch(`/users/${id}`, { method: 'GET' });
    return res.json();
  },
  async updateUser(id: string, data: any) {
    const res = await authenticatedFetch(`/users/${id}`, {
      method: 'PUT',
      body: JSON.stringify(data)
    });
    return res.json();
  },
  async deleteUser(id: string) {
    const res = await authenticatedFetch(`/users/${id}`, { method: 'DELETE' });
    return res.json();
  }
}; 