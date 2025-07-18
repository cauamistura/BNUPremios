import type { Buyer } from "./Buyers";

export interface Reward {
    id: string;
    owner_id: string;
    name: string;
    description: string;
    image: string;
    draw_date: string;
    completed: boolean;
    created_at: string;
    updated_at: string;
    images: string[];
    price: number;
    min_quota: number;
}

// Objeto para detalhes do prêmio com compradores
export interface RewardDetails extends Reward {
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