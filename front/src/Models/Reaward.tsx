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
    winner_user?: {
        id: string;
        name: string;
        email: string;
        role: string;
        active: boolean;
        created_at: string;
        updated_at: string;
    };
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

// Interface para a resposta do sorteio
export interface DrawResponse {
    drawn_at: string;
    message: string;
    reward_id: string;
    winner_number: number;
    winner_user: {
        active: boolean;
        created_at: string;
        email: string;
        id: string;
        name: string;
        role: string;
        updated_at: string;
    };
}