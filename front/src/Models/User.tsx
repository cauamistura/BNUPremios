export interface User {
    id: string;
    name: string;
    email: string;
    phone?: string;
    avatar?: string;
    joinDate?: string;
}

export interface Purchase {
    id: number;
    rewardId: string;
    rewardName: string;
    rewardImage: string;
    numbers: number[];
    purchaseDate: string;
    totalAmount: number;
    status: 'active' | 'completed' | 'cancelled';
}

export interface PurchasesResponse {
    purchases: Purchase[];
    pagination: {
        page: number;
        limit: number;
        total: number;
        pages: number;
        has_next: boolean;
        has_prev: boolean;
    };
} 