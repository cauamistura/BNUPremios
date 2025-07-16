import React from 'react';
import type { Reward } from '../../Models/Reaward';
import RewardCard from '../RewardCard';
import './index.css';

interface RewardCardListProps {
    rewards: Reward[];
}

const RewardCardList: React.FC<RewardCardListProps> = ({ rewards }) => {
    return (
        <div className="reward-list">
            {rewards.map(reward => (
                <RewardCard 
                    key={reward.id} 
                    reward={reward} 
                />
            ))}
        </div>
    );
};

export default RewardCardList;
