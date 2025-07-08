import React from 'react';
import { useAuth } from '../../hooks/useAuth';
import LoginModal from '../LoginModal';

interface ProtectedRouteProps {
    children: React.ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
    const { isAuthenticated, loading } = useAuth();

    console.log('ProtectedRoute - isAuthenticated:', isAuthenticated, 'loading:', loading);

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
            return (
                <div style={{ 
                    display: 'flex', 
                    justifyContent: 'center', 
                    alignItems: 'center', 
                    height: '50vh',
                    flexDirection: 'column',
                    gap: '20px'
                }}>
                    <h2>Área Restrita</h2>
                    <p>Você precisa estar logado para acessar esta página.</p>
                    <LoginModal 
                        isOpen={true}
                        onClose={() => window.history.back()}
                        onSuccess={() => window.location.reload()}
                    />
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