import React, { createContext, useState, useEffect } from 'react';
import type { ReactNode } from 'react';
import type { User } from '../Models/User';

interface AuthContextType {
    user: User | null;
    isAuthenticated: boolean;
    login: (email: string, password: string) => Promise<boolean>;
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
            // Simular uma chamada de API
            await new Promise(resolve => setTimeout(resolve, 1000));
            
            // Mock de validação (em produção, isso seria uma chamada real para a API)
            if (email === 'joao.silva@email.com' && password === '123456') {
                const mockUser: User = {
                    id: 1,
                    name: 'João Silva',
                    email: 'joao.silva@email.com',
                    phone: '(11) 99999-9999',
                    avatar: 'https://via.placeholder.com/150/4A90E2/FFFFFF?text=JS',
                    joinDate: '2024-01-15'
                };
                
                setUser(mockUser);
                localStorage.setItem('user', JSON.stringify(mockUser));
                setLoading(false);
                return true;
            } else {
                setLoading(false);
                return false;
            }
        } catch (error) {
            console.error('Erro no login:', error);
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
        logout,
        loading
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};

 