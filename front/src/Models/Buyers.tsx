export interface User {
    id: string;
    name: string;
    email: string;
    role: string;
    active: boolean;
    created_at: string;
    updated_at: string;
}

export interface Buyer {
    user: User;
    total_numbers: number;
}
