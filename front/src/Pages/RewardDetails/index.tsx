import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import type { RewardDetails } from '../../Models/Reaward';
import { rewardsService } from '../../services/rewardsService';
import './index.css';
import ImageCarousel from '../../Components/ImageCarousel';
import TopBuyers from '../../Components/TopBuyers';
import { formatDate } from '../../utils/formatDate';
import QuotaSelector from '../../Components/QuotaSelector';

export default function RewardDetails() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const [reward, setReward] = useState<RewardDetails | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchRewardDetails = async () => {
            if (!id) return;
            
            try {
                setLoading(true);
                setError(null);
                const data = await rewardsService.getRewardDetails(id);
                setReward(data);
            } catch (err) {
                setError('Erro ao carregar os detalhes do prêmio. Tente novamente.');
                console.error('Erro ao buscar detalhes do prêmio:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchRewardDetails();
    }, [id]);

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
                            <QuotaSelector price={reward.price} minQuota={reward.min_quota} />
                        </div>
                        <div className="reward-details-actions">
                            <button className="participate-button">
                                Participar do Sorteio
                            </button>
                        </div>
                    </div>
                    
                    {reward.buyers && reward.buyers.length > 0 && (
                        <TopBuyers buyers={reward.buyers} />
                    )}
                </div>
            </div>
        </div>
    );
} 