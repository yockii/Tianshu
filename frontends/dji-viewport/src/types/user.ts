export interface LoginResponse {
    token: string;
    user: User;
    permission: string[];
}

export interface User {
    id: string;
    tenantId: string;
    username: string;
    email: string;
    phone: string;
    createdAt: string;
    updatedAt: string;
}