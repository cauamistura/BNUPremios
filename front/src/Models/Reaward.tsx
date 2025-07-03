import type { Buyer } from "./Buyers";

export interface Reward {
    id: number;
    name: string;
    description: string;
    image: string;    
    drawDate: string;
    completed: boolean;
}

// Novo objeto para detalhes do prêmio
export interface RewardDetails extends Reward {
    images: string[];
    price: number;
    minQuota: number;
    buyers: Buyer[];
}