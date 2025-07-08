export interface User {
    id: number;
    name: string;
    email: string;
    phone: string;
    avatar?: string;
    joinDate: string;
}

export interface Purchase {
    id: number;
    rewardId: number;
    rewardName: string;
    rewardImage: string;
    numbers: number[];
    purchaseDate: string;
    totalAmount: number;
    status: 'active' | 'completed' | 'cancelled';
} 