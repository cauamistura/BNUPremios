import type { Reward, RewardsResponse, RewardDetails } from '../Models/Reaward';

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
    }
}; 