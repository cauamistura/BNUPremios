import type { PurchasesResponse } from '../Models/User';
import { authenticatedFetch } from './apiUtils';

export const purchasesService = {
    async getUserPurchases(userId: string): Promise<PurchasesResponse> {
        try {
            const response = await authenticatedFetch(`/purchases/user/${userId}`, { method: 'GET' });
            const data: PurchasesResponse = await response.json();
            
            // Garantir que purchases seja sempre um array válido
            return {
                purchases: Array.isArray(data.purchases) ? data.purchases : [],
                pagination: data.pagination || {
                    page: 1,
                    limit: 10,
                    total: 0,
                    pages: 1,
                    has_next: false,
                    has_prev: false
                }
            };
        } catch (error) {
            console.error('Erro ao buscar compras do usuário:', error);
            // Retornar estrutura válida em caso de erro
            return {
                purchases: [],
                pagination: {
                    page: 1,
                    limit: 10,
                    total: 0,
                    pages: 1,
                    has_next: false,
                    has_prev: false
                }
            };
        }
    }
}; 