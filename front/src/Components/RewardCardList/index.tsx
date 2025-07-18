import React from 'react';
import type { Reward } from '../../Models/Reaward';
import RewardCard from '../RewardCard';
import './index.css';

interface RewardCardListProps {
    rewards: Reward[];
    routeItem: String;
    onEdit?: (reward: Reward) => void;
    onDelete?: (reward: Reward) => void;
}

const RewardCardList: React.FC<RewardCardListProps> = ({ rewards, routeItem, onEdit, onDelete }) => {
    // Verificação de segurança
    if (!rewards || !Array.isArray(rewards) || rewards.length === 0) {
        return null;
    }

    return (
        <div className="reward-list">
            {rewards.map(reward => (
                <RewardCard 
                    key={reward.id} 
                    reward={reward} 
                    routeItem={routeItem}
                    onEdit={onEdit}
                    onDelete={onDelete}
                />
            ))}
        </div>
    );
};

export default RewardCardList;
