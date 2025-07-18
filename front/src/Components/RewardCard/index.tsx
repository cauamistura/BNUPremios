import React from 'react';
import { useNavigate } from 'react-router-dom';
import type { Reward } from '../../Models/Reaward';
import './index.css';

interface RewardCardProps {
    reward: Reward;
    routeItem: String;
    onEdit?: (reward: Reward) => void;
    onDelete?: (reward: Reward) => void;
    onDraw?: (reward: Reward) => void;
}

const RewardCard: React.FC<RewardCardProps> = ({ reward, routeItem, onEdit, onDelete, onDraw}) => {
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
        navigate(`/${routeItem}/${reward.id}`);
    };

    const handleEditClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (onEdit) {
            onEdit(reward);
        }
    };

    const handleDeleteClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (onDelete) {
            onDelete(reward);
        }
    };

    const handleDrawClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (onDraw) {
            onDraw(reward);
        }
    };

    return (
        <div className="reward-card" onClick={handleCardClick}>
            {onEdit && (
                <button 
                    className="edit-button"
                    onClick={handleEditClick}
                    title="Editar prêmio"
                >
                    ✏️
                </button>
            )}
            {onDelete && (
                <button 
                    className="delete-button"
                    onClick={handleDeleteClick}
                    title="Excluir prêmio"
                >
                    🗑️
                </button>
            )}
            {onDraw && !reward.completed && (
                <button 
                    className="draw-button"
                    onClick={handleDrawClick}
                    title="Realizar sorteio"
                >
                    🎲
                </button>
            )}
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
