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
            return data;
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
    async createReward(data: any) {
        const res = await authenticatedFetch('/rewards', {
            method: 'POST',
            body: JSON.stringify(data)
        });
        return res.json();
    },
    async updateReward(id: string, data: any) {
        const res = await authenticatedFetch(`/rewards/${id}`, {
            method: 'PUT',
            body: JSON.stringify(data)
        });
        return res.json();
    },
    async deleteReward(id: string) {
        const res = await authenticatedFetch(`/rewards/${id}`, { method: 'DELETE' });
        return res.json();
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
    }
}; 