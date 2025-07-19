import React, { useState, useEffect } from 'react';
import type { Purchase } from '../../Models/User';
import { useAuth } from '../../hooks/useAuth';
import { purchasesService } from '../../services/purchasesService';
import { useToastContext } from '../../contexts/ToastContext';
import './index.css';

const Profile: React.FC = () => {
    const { user: authUser, logout, loading: authLoading } = useAuth();
    const { showError } = useToastContext();
    const [purchases, setPurchases] = useState<Purchase[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUserPurchases = async () => {
            if (!authUser?.id) {
                setLoading(false);
                return;
            }

            try {
                setLoading(true);
                const response = await purchasesService.getUserPurchases(authUser.id);
                // Garantir que purchases seja sempre um array válido
                setPurchases(response.purchases || []);
            } catch (err) {
                const errorMessage = 'Erro ao carregar suas compras. Tente novamente.';
                showError(errorMessage);
                console.error('Erro ao buscar compras do usuário:', err);
                // Em caso de erro, garantir que purchases seja um array vazio
                setPurchases([]);
            } finally {
                setLoading(false);
            }
        };

        // Só busca as compras se a autenticação não estiver carregando e o usuário existir
        if (!authLoading && authUser?.id) {
            fetchUserPurchases();
        } else if (!authLoading && !authUser) {
            setLoading(false);
        }
    }, [authUser?.id, authLoading]);

    const getStatusColor = (status: string) => {
        switch (status) {
            case 'active':
                return '#4CAF50';
            case 'completed':
                return '#2196F3';
            case 'cancelled':
                return '#F44336';
            default:
                return '#757575';
        }
    };

    const getStatusText = (status: string) => {
        switch (status) {
            case 'active':
                return 'Ativo';
            case 'completed':
                return 'Concluído';
            case 'cancelled':
                return 'Cancelado';
            default:
                return status;
        }
    };

    const formatDate = (dateString: string) => {
        return new Date(dateString).toLocaleDateString('pt-BR');
    };

    if (authLoading || loading) {
        return (
            <div className="profile-container">
                <div className="profile-loading-container">
                    <div className="profile-loading-spinner"></div>
                    <p>{authLoading ? 'Carregando dados do usuário...' : 'Carregando suas compras...'}</p>
                </div>
            </div>
        );
    }

    if (!authUser) {
        return (
            <div className="profile-container">
                <div className="profile-error-container">
                    <p className="profile-error-message">Erro ao carregar dados do usuário</p>
                </div>
            </div>
        );
    }

    return (
        <div className="profile-container">            

            {/* Seção de informações do usuário */}
            <div className="profile-user-info-section">
                
                <div className="profile-user-details">
                    <h2>{authUser.name}</h2>
                    <div className="profile-user-info-grid">
                        <div className="profile-info-item">
                            <span className="profile-label">Email:</span>
                            <span className="profile-value">{authUser.email}</span>
                        </div>
                        {authUser.phone && (
                            <div className="profile-info-item">
                                <span className="profile-label">Telefone:</span>
                                <span className="profile-value">{authUser.phone}</span>
                            </div>
                        )}
                        {authUser.joinDate && (
                            <div className="profile-info-item">
                                <span className="profile-label">Membro desde:</span>
                                <span className="profile-value">{formatDate(authUser.joinDate)}</span>
                            </div>
                        )}
                        <div className="profile-info-item">
                            <button 
                                onClick={logout}
                                className="profile-logout-btn"
                            >
                                Sair da Conta
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            {/* Seção de compras */}
            <div className="profile-purchases-section">
                <h3>Minhas Compras</h3>
                {!purchases || purchases.length === 0 ? (
                    <div className="profile-no-purchases">
                        <p>Você ainda não fez nenhuma compra.</p>
                    </div>
                ) : (
                    <div className="profile-purchases-grid">
                        {purchases.map((purchase) => (
                            <div key={purchase.id} className="profile-purchase-card">
                                <div className="profile-purchase-header">
                                    <img 
                                        src={purchase.rewardImage} 
                                        alt={purchase.rewardName}
                                        className="profile-reward-image"
                                    />
                                    <div className="profile-purchase-info">
                                        <h4>{purchase.rewardName}</h4>
                                        <p className="profile-purchase-date">
                                            Comprado em: {formatDate(purchase.purchaseDate)}
                                        </p>
                                        <p className="profile-purchase-amount">
                                            Valor: R$ {purchase.totalAmount.toFixed(2)}
                                        </p>
                                        <span 
                                            className="profile-status-badge"
                                            style={{ backgroundColor: getStatusColor(purchase.status) }}
                                        >
                                            {getStatusText(purchase.status)}
                                        </span>
                                    </div>
                                </div>
                                <div className="profile-numbers-section">
                                    <h5>Números comprados:</h5>
                                    <div className="profile-numbers-grid">
                                        {purchase.numbers.map((number, index) => (
                                            <span key={index} className="profile-number-badge">
                                                {number.toString().padStart(2, '0')}
                                            </span>
                                        ))}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
};

export default Profile; 