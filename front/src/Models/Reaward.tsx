export interface Reward {
    id: number;
    name: string;
    description: string;
    image: string;
    drawDate: string; // Data no formato ISO "YYYY-MM-DD"
    completed: boolean;
}

// Tipo para array de prêmios
export type RewardList = Reward[];

