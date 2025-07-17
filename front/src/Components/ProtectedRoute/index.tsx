import React, { useEffect } from 'react';
import { useAuth } from '../../hooks/useAuth';
import { useRedirectToLogin } from '../../hooks/useRedirectToLogin';

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
    const { isAuthenticated, loading } = useAuth();
    const { redirectToLogin } = useRedirectToLogin();

    console.log('ProtectedRoute - isAuthenticated:', isAuthenticated, 'loading:', loading);

    useEffect(() => {
        if (!loading && !isAuthenticated) {
            redirectToLogin();
        }
    }, [loading, isAuthenticated, redirectToLogin]);

    try {
        if (loading) {
            return (
                <div style={{ 
                    display: 'flex', 
                    justifyContent: 'center', 
                    alignItems: 'center', 
                    height: '50vh',
                    fontSize: '1.2rem',
                    color: '#666'
                }}>
                    Verificando autenticação...
                </div>
            );
        }

        if (!isAuthenticated) {
            // Retorna um loading enquanto redireciona
            return (
                <div style={{ 
                    display: 'flex', 
                    justifyContent: 'center', 
                    alignItems: 'center', 
                    height: '50vh',
                    fontSize: '1.2rem',
                    color: '#666'
                }}>
                    Redirecionando para login...
                </div>
            );
        }

        return <>{children}</>;
    } catch (error) {
        console.error('Erro no ProtectedRoute:', error);
        return (
            <div style={{ 
                display: 'flex', 
                justifyContent: 'center', 
                alignItems: 'center', 
                height: '50vh',
                fontSize: '1.2rem',
                color: '#f44336'
            }}>
                Erro ao carregar a página: {error instanceof Error ? error.message : 'Erro desconhecido'}
            </div>
        );
    }
};

export default ProtectedRoute; 