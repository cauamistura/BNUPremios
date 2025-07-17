import type { PurchasesResponse } from '../Models/User';
import { authenticatedFetch } from './apiUtils';

export const purchasesService = {
    async getUserPurchases(userId: string): Promise<PurchasesResponse> {
        const response = await authenticatedFetch(`/purchases/user/${userId}`, { method: 'GET' });
        return response.json();
    }
}; 