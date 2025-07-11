import React, { createContext, useState, useEffect } from 'react';
import type { ReactNode } from 'react';
import type { User } from '../Models/User';
import type { RegisterPayload } from '../services/authService';
import { login as apiLogin, register as apiRegister } from '../services/authService';

interface AuthContextType {
    user: User | null;
    isAuthenticated: boolean;
    login: (email: string, password: string) => Promise<boolean>;
    register: (payload: RegisterPayload) => Promise<boolean>;
    logout: () => void;
    loading: boolean;
}

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
    children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    // Verificar se há um usuário salvo no localStorage ao inicializar
    useEffect(() => {
        const savedUser = localStorage.getItem('user');
        if (savedUser) {
            try {
                setUser(JSON.parse(savedUser));
            } catch (error) {
                console.error('Erro ao carregar usuário do localStorage:', error);
                localStorage.removeItem('user');
            }
        }
        setLoading(false);
    }, []);

    const login = async (email: string, password: string): Promise<boolean> => {
        setLoading(true);
        try {
            const data = await apiLogin({ email, password });
            setUser(data.user);
            localStorage.setItem('user', JSON.stringify(data.user));
            localStorage.setItem('token', data.token);
            setLoading(false);
            return true;
        } catch {
            setLoading(false);
            return false;
        }
    };

    const register = async (payload: RegisterPayload): Promise<boolean> => {
        setLoading(true);
        try {
            await apiRegister(payload);
            setLoading(false);
            return true;
        } catch {
            setLoading(false);
            return false;
        }
    };

    const logout = () => {
        setUser(null);
        localStorage.removeItem('user');
    };

    const value: AuthContextType = {
        user,
        isAuthenticated: !!user,
        login,
        register,
        logout,
        loading
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};

 