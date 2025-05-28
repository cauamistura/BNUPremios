import React from 'react';
import type { Reward } from '../../Models/Reaward';
import './index.css';

interface RewardCardProps {
    reward: Reward;
}

const RewardCard: React.FC<RewardCardProps> = ({ reward }) => {
    const formatDate = (date: string) => {
        return new Date(date).toLocaleDateString('pt-BR', {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
        });
    };    

    return (
        <div className="reward-card">
            <div className="reward-image">
                <img src={reward.image} alt={reward.name} />
            </div>
            <div className="reward-info">
                <h2 className="reward-title">{reward.name}</h2>
                <p className="reward-description">{reward.description}</p>
                <div className="reward-status">
                    <p className="reward-date">Data do Sorteio: {formatDate(reward.drawDate)}</p>
                    <span className={`status-badge ${reward.completed ? 'completed' : 'pending'}`}>
                        {reward.completed ? 'Sorteado' : 'Disponível'}
                    </span>
                </div>
            </div>
        </div>
    );
};

export default RewardCard;
