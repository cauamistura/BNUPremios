import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import type { RewardDetails } from '../../Models/Reaward';
import { rewardsService } from '../../services/rewardsService';
import { useAuth } from '../../hooks/useAuth';
import { useToast } from '../../hooks/useToast';
import { useRedirectToLogin } from '../../hooks/useRedirectToLogin';
import './index.css';
import ImageCarousel from '../../Components/ImageCarousel';
import TopBuyers from '../../Components/TopBuyers';
import ToastContainer from '../../Components/ToastContainer';
import { formatDate } from '../../utils/formatDate';
import QuotaSelector from '../../Components/QuotaSelector';

export default function RewardDetails() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const { user: authUser } = useAuth();
    const { redirectToLogin } = useRedirectToLogin();
    const { toasts, removeToast, showSuccess, showError, showWarning } = useToast();
    const [reward, setReward] = useState<RewardDetails | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [selectedQuantity, setSelectedQuantity] = useState(1);
    const [buying, setBuying] = useState(false);
    const [showConfirmation, setShowConfirmation] = useState(false);

    useEffect(() => {
        const fetchRewardDetails = async () => {
            if (!id) return;
            
            try {
                setLoading(true);
                setError(null);
                const data = await rewardsService.getRewardDetails(id);
                setReward(data);
                setSelectedQuantity(data.min_quota);
            } catch (err) {
                setError('Erro ao carregar os detalhes do prêmio. Tente novamente.');
                console.error('Erro ao buscar detalhes do prêmio:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchRewardDetails();
    }, [id]);

    const handleQuantityChange = (quantity: number) => {
        setSelectedQuantity(quantity);
    };

    const handleParticipateClick = () => {
        if (!authUser) {
            showWarning('Você precisa estar logado para participar do sorteio. Redirecionando para a página de login...');
            // Pequeno delay para o usuário ver a mensagem antes do redirecionamento
            setTimeout(() => {
                redirectToLogin();
            }, 2000);
            return;
        }

        if (selectedQuantity < reward!.min_quota) {
            showError(`A quantidade mínima é ${reward!.min_quota} números.`);
            return;
        }

        setShowConfirmation(true);
    };

    const handleConfirmPurchase = async () => {
        if (!reward || !authUser) return;

        try {
            setBuying(true);
            setError(null);
            
            const result = await rewardsService.buyNumbers(reward.id, authUser.id, selectedQuantity);
            
            showSuccess(`Compra realizada com sucesso! Números: ${result.numbers.join(', ')}`);
            
            // Recarregar os detalhes da recompensa para atualizar a lista de compradores
            const updatedReward = await rewardsService.getRewardDetails(reward.id);
            setReward(updatedReward);
            
            setShowConfirmation(false);
        } catch (err: any) {
            showError(err.message || 'Erro ao realizar a compra. Tente novamente.');
            console.error('Erro ao comprar números:', err);
        } finally {
            setBuying(false);
        }
    };

    const handleCancelPurchase = () => {
        setShowConfirmation(false);
    };

    if (loading) {
        return (
            <div className="reward-details-loading-container">
                <div className="reward-details-loading-spinner"></div>
                <p>Carregando detalhes do prêmio...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="reward-details-error-container">
                <p className="reward-details-error-message">{error}</p>
                <button 
                    className="reward-details-error-reload-btn" 
                    onClick={() => window.location.reload()}
                >
                    Tentar novamente
                </button>
            </div>
        );
    }

    if (!reward) {
        return (
            <div className="reward-details-container">
                <div className="reward-details-content">
                    <h1>Prêmio não encontrado</h1>
                    <button onClick={() => navigate('/')} className="back-button">
                        Voltar para Home
                    </button>
                </div>
            </div>
        );
    }
        
    return (
        <div className="reward-details-container">
            <div className="reward-details-content">
                <div className="reward-details-card">
                    <div className="reward-details-carousel">
                        <ImageCarousel images={reward.images} altPrefix={reward.name} />
                    </div>                                        
                    
                    <div className="reward-details-info">
                        <h1 className="reward-details-title">{reward.name}</h1>
                        <p className="reward-details-description">{reward.description}</p>
                        <hr className="reward-divider" />
                        <div className="reward-details-meta">
                            <div className="reward-details-date">
                                <strong>Data do Sorteio:</strong> {formatDate(reward.draw_date)}
                                <span className={`status-badge ${reward.completed ? 'completed' : 'pending'}`}> 
                                    {reward.completed ? 'Sorteado' : 'Disponível'}
                                </span>
                            </div>                            
                        </div>
                        <div className="reward-price-box">
                            <div className="reward-price-sober">
                                Preço por cota: <span>R$ {reward.price.toLocaleString('pt-BR', { minimumFractionDigits: 2 })}</span>
                            </div>
                            <QuotaSelector 
                                price={reward.price} 
                                minQuota={reward.min_quota} 
                                onChange={handleQuantityChange}
                            />
                        </div>
                        <div className="reward-details-actions">
                            <button 
                                className="participate-button"
                                onClick={handleParticipateClick}
                                disabled={buying}
                            >
                                {buying ? 'Processando...' : 'Participar do Sorteio'}
                            </button>
                        </div>
                    </div>
                    
                    {reward.buyers && Array.isArray(reward.buyers) && reward.buyers.length > 0 && (
                        <TopBuyers buyers={reward.buyers} />
                    )}
                </div>
            </div>

            {/* Modal de confirmação */}
            {showConfirmation && (
                <div className="reward-confirmation-overlay">
                    <div className="reward-confirmation-modal">
                        <h3>Confirmar Compra</h3>
                        <div className="reward-confirmation-details">
                            <p><strong>Prêmio:</strong> {reward.name}</p>
                            <p><strong>Quantidade:</strong> {selectedQuantity} números</p>
                            <p><strong>Valor total:</strong> R$ {(selectedQuantity * reward.price).toLocaleString('pt-BR', { minimumFractionDigits: 2 })}</p>
                        </div>
                        <div className="reward-confirmation-actions">
                            <button 
                                className="reward-confirm-btn"
                                onClick={handleConfirmPurchase}
                                disabled={buying}
                            >
                                {buying ? 'Processando...' : 'Confirmar'}
                            </button>
                            <button 
                                className="reward-cancel-btn"
                                onClick={handleCancelPurchase}
                                disabled={buying}
                            >
                                Cancelar
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Toast Container */}
            <ToastContainer toasts={toasts} onRemove={removeToast} />
        </div>
    );
} 