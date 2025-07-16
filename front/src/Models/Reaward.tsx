import type { Buyer } from "./Buyers";

export interface Reward {
    id: string;
    name: string;
    description: string;
    image: string;    
    draw_date: string;
    completed: boolean;
    created_at: string;
    updated_at: string;
}

// Novo objeto para detalhes do prêmio
export interface RewardDetails extends Reward {
    images: string[];
    price: number;
    min_quota: number;
    buyers: Buyer[] | null;
}

export interface Pagination {
    page: number;
    limit: number;
    total: number;
    pages: number;
    has_next: boolean;
    has_prev: boolean;
}

export interface RewardsResponse {
    rewards: Reward[];
    pagination: Pagination;
}