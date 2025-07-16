import React from 'react';
import { useNavigate } from 'react-router-dom';
import type { Reward } from '../../Models/Reaward';
import './index.css';

interface RewardCardProps {
    reward: Reward;
}

const RewardCard: React.FC<RewardCardProps> = ({ reward }) => {
    const navigate = useNavigate();
    
    const formatDate = (date: string) => {
        return new Date(date).toLocaleDateString('pt-BR', {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
        });
    };

    const handleCardClick = () => {
        navigate(`/premio/${reward.id}`);
    };

    return (
        <div className="reward-card" onClick={handleCardClick}>
            <div className="reward-image">
                <img src={reward.image} alt={reward.name} />
            </div>
            <div className="reward-info">
                <h2 className="reward-title">{reward.name}</h2>
                <p className="reward-description">{reward.description}</p>
                <div className="reward-status">
                    <p className="reward-date">Data do Sorteio: {formatDate(reward.draw_date)}</p>
                    <span className={`reward-card-status-badge ${reward.completed ? 'completed' : 'pending'}`}>
                        {reward.completed ? 'Sorteado' : 'Disponível'}
                    </span>
                </div>
            </div>
        </div>
    );
};

export default RewardCard;
