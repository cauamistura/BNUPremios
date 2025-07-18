import type { Reward, RewardsResponse, RewardDetails } from '../Models/Reaward';
import { authenticatedFetch } from './apiUtils';

const API_BASE_URL = '/api/v1';

export const rewardsService = {
    async getRewards(): Promise<RewardsResponse> {
        try {
            const response = await fetch(`${API_BASE_URL}/rewards`);
            if (!response.ok) {
                throw new Error(`Erro na requisição: ${response.status}`);
            }
            const data: RewardsResponse = await response.json();
            
            // Garantir que rewards seja sempre um array válido
            return {
                rewards: Array.isArray(data.rewards) ? data.rewards : [],
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
            console.error('Erro ao buscar recompensas:', error);
            throw error;
        }
    },
    async getRewardById(id: string): Promise<Reward> {
        try {
            const response = await fetch(`${API_BASE_URL}/rewards/${id}`);
            if (!response.ok) {
                throw new Error(`Erro na requisição: ${response.status}`);
            }
            const data: Reward = await response.json();
            return data;
        } catch (error) {
            console.error('Erro ao buscar recompensa:', error);
            throw error;
        }
    },
    async getRewardDetails(id: string): Promise<RewardDetails> {
        try {
            const response = await fetch(`${API_BASE_URL}/rewards/${id}/details`);
            if (!response.ok) {
                throw new Error(`Erro na requisição: ${response.status}`);
            }
            const data: RewardDetails = await response.json();
            return data;
        } catch (error) {
            console.error('Erro ao buscar detalhes da recompensa:', error);
            throw error;
        }
    },
    async buyNumbers(rewardId: string, userId: string, quantity: number): Promise<{message: string, numbers: number[], quantity: number}> {
        const response = await authenticatedFetch(`/rewards/${rewardId}/buyers/${userId}`, {
            method: 'POST',
            body: JSON.stringify({ quantity })
        });
        return response.json();
    },
    // Métodos protegidos
    async deleteReward(id: string): Promise<{ message: string }> {
        try {
            const response = await authenticatedFetch(`/rewards/${id}`, { method: 'DELETE' });
            
            // Verifica se a resposta tem conteúdo antes de tentar fazer parse
            const text = await response.text();
            
            if (text) {
                return JSON.parse(text);
            } else {
                // Se a resposta está vazia, retorna uma mensagem de sucesso
                return { message: 'Prêmio excluído com sucesso' };
            }
        } catch (error) {
            console.error('Erro ao deletar prêmio:', error);
            throw error;
        }
    },
    async addBuyer(rewardId: string, userId: string, quantity: number) {
        const res = await authenticatedFetch(`/rewards/${rewardId}/buyers/${userId}`, {
            method: 'POST',
            body: JSON.stringify({ quantity })
        });
        return res.json();
    },
    async removeBuyer(rewardId: string, userId: string) {
        const res = await authenticatedFetch(`/rewards/${rewardId}/buyers/${userId}`, {
            method: 'DELETE'
        });
        return res.json();
    },
    async getUserNumbers(rewardId: string, userId: string) {
        const res = await authenticatedFetch(`/rewards/${rewardId}/buyers/${userId}/numbers`, {
            method: 'GET'
        });
        return res.json();
    },
    async getMyRewards(): Promise<RewardsResponse> {
        try {
            const response = await authenticatedFetch('/rewards/mine');
            if (!response.ok) {
                throw new Error(`Erro na requisição: ${response.status}`);
            }
            const data: RewardsResponse = await response.json();
            
            // Garantir que rewards seja sempre um array válido
            return {
                rewards: Array.isArray(data.rewards) ? data.rewards : [],
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
            console.error('Erro ao buscar meus prêmios:', error);
            throw error;
        }
    },
    async createReward(data: {
        name: string;
        description: string;
        image: string;
        images: string[];
        draw_date: string;
        min_quota: number;
        price: number;        
    }): Promise<Reward> {
        try {
            const response = await authenticatedFetch('/rewards', {
                method: 'POST',
                body: JSON.stringify(data)
            });
            return response.json();
        } catch (error) {
            console.error('Erro ao criar prêmio:', error);
            throw error;
        }
    },
    async updateReward(id: string, data: {
        name: string;
        description: string;
        image: string;
        images: string[];
        draw_date: string;
        min_quota: number;
        price: number;
    }): Promise<Reward> {
        console.log('updateReward - dados recebidos:', data);
        try {
            const response = await authenticatedFetch(`/rewards/${id}`, {
                method: 'PUT',
                body: JSON.stringify(data)
            });
            return response.json();
        } catch (error) {
            console.error('Erro ao atualizar prêmio:', error);
            throw error;
        }
    }
}; 