import { useParams, useNavigate } from 'react-router-dom';
import type { RewardDetails } from '../../Models/Reaward';
import rewardsMock from '../../assets/Mocks/RewardsDetails.json';
import './index.css';
import ImageCarousel from '../../Components/ImageCarousel';
import TopBuyers from '../../Components/TopBuyers';
import { formatDate } from '../../utils/formatDate';
import QuotaSelector from '../../Components/QuotaSelector';

export default function RewardDetails() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    
    // Encontrar o prêmio pelo ID
    const reward: RewardDetails | undefined = rewardsMock.find(r => r.id === Number(id));
    
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
                                <strong>Data do Sorteio:</strong> {formatDate(reward.drawDate)}
                                <span className={`status-badge ${reward.completed ? 'completed' : 'pending'}`}> 
                                    {reward.completed ? 'Sorteado' : 'Disponível'}
                                </span>
                            </div>                            
                        </div>
                        <div className="reward-price-box">
                            <div className="reward-price-sober">
                                Preço por cota: <span>R$ {reward.price.toLocaleString('pt-BR', { minimumFractionDigits: 2 })}</span>
                            </div>
                            <QuotaSelector price={reward.price} minQuota={reward.minQuota} />
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