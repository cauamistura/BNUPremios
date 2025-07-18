import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import RewardCardList from "../../Components/RewardCardList";
import { rewardsService } from "../../services/rewardsService";
import type { Reward, DrawResponse } from "../../Models/Reaward";
import { useToastContext } from "../../contexts/ToastContext";
import "./index.css";

export default function MyRewards() {
    const navigate = useNavigate();
    const { showError, showSuccess } = useToastContext();
    const [rewards, setRewards] = useState<Reward[]>([]);
    const [loading, setLoading] = useState(true);
    const [deletingReward, setDeletingReward] = useState<Reward | null>(null);
    const [showDeleteConfirmation, setShowDeleteConfirmation] = useState(false);
    const [deleteLoading, setDeleteLoading] = useState(false);
    
    // Estados para o sorteio
    const [drawingReward, setDrawingReward] = useState<Reward | null>(null);
    const [showDrawConfirmation, setShowDrawConfirmation] = useState(false);
    const [drawLoading, setDrawLoading] = useState(false);
    const [drawResult, setDrawResult] = useState<DrawResponse | null>(null);
    const [showDrawResult, setShowDrawResult] = useState(false);

    const fetchMyRewards = async () => {
        try {
            setLoading(true);
            const response = await rewardsService.getMyRewards();
            setRewards(response.rewards || []);
        } catch (err) {
            const errorMessage = 'Erro ao carregar seus prêmios. Tente novamente.';
            showError(errorMessage);
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
            
            // Mostra mensagem de sucesso
            showSuccess('Prêmio excluído com sucesso!');
            console.log('Prêmio excluído:', result.message);
        } catch (err) {
            console.error('Erro ao deletar prêmio:', err);
            
            // Extrai a mensagem de erro
            let errorMessage = 'Erro ao deletar prêmio. Tente novamente.';
            if (err instanceof Error) {
                errorMessage = err.message;
            }
            
            showError(errorMessage);
        } finally {
            setDeleteLoading(false);
        }
    };

    const handleCancelDelete = () => {
        setShowDeleteConfirmation(false);
        setDeletingReward(null);
    };

    const handleDrawReward = (reward: Reward) => {
        setDrawingReward(reward);
        setShowDrawConfirmation(true);
    };

    const handleConfirmDraw = async () => {
        if (!drawingReward) return;

        try {
            setDrawLoading(true);
            const result = await rewardsService.performDraw(drawingReward.id);
            setShowDrawConfirmation(false);
            setDrawingReward(null);
            setDrawResult(result);
            setShowDrawResult(true);
            fetchMyRewards(); // Recarrega a lista para atualizar o status
            
            // Mostra mensagem de sucesso
            showSuccess('Sorteio realizado com sucesso!');
        } catch (err) {
            console.error('Erro ao realizar sorteio:', err);
            
            // Extrai a mensagem de erro
            let errorMessage = 'Erro ao realizar sorteio. Tente novamente.';
            if (err instanceof Error) {
                errorMessage = err.message;
            }
            
            showError(errorMessage);
            setShowDrawConfirmation(false);
            setDrawingReward(null);
        } finally {
            setDrawLoading(false);
        }
    };

    const handleCancelDraw = () => {
        setShowDrawConfirmation(false);
        setDrawingReward(null);
    };

    const handleCloseDrawResult = () => {
        setShowDrawResult(false);
        setDrawResult(null);
    };

    const formatDate = (date: string) => {
        return new Date(date).toLocaleDateString('pt-BR', {
            day: '2-digit',
            month: '2-digit',
            year: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
        });
    };

    if (loading) {
        return (
            <div className="my-rewards-loading-container">
                <div className="my-rewards-loading-spinner"></div>
                <p>Carregando seus prêmios...</p>
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
                    onDraw={handleDrawReward}
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

            {/* Modal de confirmação de sorteio */}
            {showDrawConfirmation && drawingReward && (
                <div className="delete-confirmation-overlay">
                    <div className="delete-confirmation-modal">
                        <h3>Confirmar Sorteio</h3>
                        <p>Tem certeza que deseja realizar o sorteio do prêmio <strong>"{drawingReward.name}"</strong>?</p>
                        <p className="delete-warning">Esta ação não pode ser desfeita.</p>
                        <div className="delete-confirmation-actions">
                            <button 
                                className="delete-confirm-btn"
                                onClick={handleConfirmDraw}
                                disabled={drawLoading}
                            >
                                {drawLoading ? 'Realizando sorteio...' : 'Realizar Sorteio'}
                            </button>
                            <button 
                                className="delete-cancel-btn"
                                onClick={handleCancelDraw}
                                disabled={drawLoading}
                            >
                                Cancelar
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Modal de resultado do sorteio */}
            {showDrawResult && drawResult && (
                <div className="delete-confirmation-overlay">
                    <div className="delete-confirmation-modal">
                        <h3>🎉 Sorteio Realizado!</h3>
                        <div className="draw-result-content">
                            <p><strong>Mensagem:</strong> {drawResult.message}</p>
                            <p><strong>Número sorteado:</strong> {drawResult.winner_number}</p>
                            <p><strong>Ganhador:</strong> {drawResult.winner_user.name}</p>
                            <p><strong>Email:</strong> {drawResult.winner_user.email}</p>
                            <p><strong>Data do sorteio:</strong> {formatDate(drawResult.drawn_at)}</p>
                        </div>
                        <div className="delete-confirmation-actions">
                            <button 
                                className="delete-confirm-btn"
                                onClick={handleCloseDrawResult}
                            >
                                Fechar
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
} 