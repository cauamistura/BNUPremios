import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import RewardCardList from "../../Components/RewardCardList";
import { rewardsService } from "../../services/rewardsService";
import type { Reward } from "../../Models/Reaward";
import "./index.css";

export default function MyRewards() {
    const navigate = useNavigate();
    const [rewards, setRewards] = useState<Reward[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [deletingReward, setDeletingReward] = useState<Reward | null>(null);
    const [showDeleteConfirmation, setShowDeleteConfirmation] = useState(false);
    const [deleteLoading, setDeleteLoading] = useState(false);

    const fetchMyRewards = async () => {
        try {
            setLoading(true);
            setError(null);
            const response = await rewardsService.getMyRewards();
            setRewards(response.rewards || []);
        } catch (err) {
            setError('Erro ao carregar seus prêmios. Tente novamente.');
            console.error('Erro ao buscar meus prêmios:', err);
            setRewards([]); // Garante que sempre temos um array
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchMyRewards();
    }, []);

    const handleCreateReward = () => {
        navigate('/MeuSorteio/novo');
    };

    const handleEditReward = (reward: Reward) => {
        navigate(`/MeuSorteio/${reward.id}`);
    };

    const handleDeleteReward = (reward: Reward) => {
        setDeletingReward(reward);
        setShowDeleteConfirmation(true);
    };

    const handleConfirmDelete = async () => {
        if (!deletingReward) return;

        try {
            setDeleteLoading(true);
            const result = await rewardsService.deleteReward(deletingReward.id);
            setShowDeleteConfirmation(false);
            setDeletingReward(null);
            fetchMyRewards(); // Recarrega a lista
            
            // Mostra mensagem de sucesso (opcional)
            console.log('Prêmio excluído:', result.message);
        } catch (err) {
            console.error('Erro ao deletar prêmio:', err);
            setError('Erro ao deletar prêmio. Tente novamente.');
        } finally {
            setDeleteLoading(false);
        }
    };

    const handleCancelDelete = () => {
        setShowDeleteConfirmation(false);
        setDeletingReward(null);
    };

    if (loading) {
        return (
            <div className="my-rewards-loading-container">
                <div className="my-rewards-loading-spinner"></div>
                <p>Carregando seus prêmios...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="my-rewards-error-container">
                <p className="my-rewards-error-message">{error}</p>
                <button className="my-rewards-error-reload-btn" onClick={() => window.location.reload()}>
                    Tentar novamente
                </button>
            </div>
        );
    }

    return (
        <div className="my-rewards-container">
            <div className="button-container-myrewards">
                <button
                    className="my-rewards-button"
                    onClick={handleCreateReward}
                >
                    Cadastrar prêmios
                </button>
            </div>

            {!rewards || rewards.length === 0 ? (
                <div className="my-rewards-empty">
                    <p>Você ainda não cadastrou prêmios.</p>
                </div>
            ) : (
                <RewardCardList
                    rewards={rewards}
                    routeItem={"MeuSorteio"}
                    onEdit={handleEditReward}
                    onDelete={handleDeleteReward}
                />
            )}

            {/* Modal de confirmação de delete */}
            {showDeleteConfirmation && deletingReward && (
                <div className="delete-confirmation-overlay">
                    <div className="delete-confirmation-modal">
                        <h3>Confirmar Exclusão</h3>
                        <p>Tem certeza que deseja excluir o prêmio <strong>"{deletingReward.name}"</strong>?</p>
                        <p className="delete-warning">Esta ação não pode ser desfeita.</p>
                        <div className="delete-confirmation-actions">
                            <button 
                                className="delete-confirm-btn"
                                onClick={handleConfirmDelete}
                                disabled={deleteLoading}
                            >
                                {deleteLoading ? 'Excluindo...' : 'Excluir'}
                            </button>
                            <button 
                                className="delete-cancel-btn"
                                onClick={handleCancelDelete}
                                disabled={deleteLoading}
                            >
                                Cancelar
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
} 