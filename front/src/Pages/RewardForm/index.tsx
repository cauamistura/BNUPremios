import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import RewardForm from "../../Components/RewardForm";
import type { RewardFormData } from "../../Components/RewardForm";
import { rewardsService } from "../../services/rewardsService";
import type { Reward } from "../../Models/Reaward";
import { useAuth } from "../../hooks/useAuth";
import { useToastContext } from "../../contexts/ToastContext";
import "./index.css";

export default function RewardFormPage() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const { user } = useAuth();
    const { showError, showSuccess } = useToastContext();
    
    const [reward, setReward] = useState<Reward | undefined>();
    const [loading, setLoading] = useState(true);
    const [formLoading, setFormLoading] = useState(false);

    useEffect(() => {
        const fetchReward = async () => {
            if (!id) {
                setLoading(false);
                return;
            }

            try {
                setLoading(true);
                const data = await rewardsService.getRewardById(id);
                setReward(data);
            } catch (err) {
                const errorMessage = 'Erro ao carregar o prêmio. Tente novamente.';
                showError(errorMessage);
                console.error('Erro ao buscar prêmio:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchReward();
    }, [id]);

    const handleCreateReward = async (formData: RewardFormData) => {
        if (!user?.id) return;
        
        try {
            setFormLoading(true);
            await rewardsService.createReward(formData);
            showSuccess('Prêmio criado com sucesso!');
            navigate('/MeusSorteios');
        } catch (err) {
            console.error('Erro ao criar prêmio:', err);
            
            // Extrai a mensagem de erro
            let errorMessage = 'Erro ao criar prêmio. Tente novamente.';
            if (err instanceof Error) {
                errorMessage = err.message;
            }
            
            showError(errorMessage);
        } finally {
            setFormLoading(false);
        }
    };

    const handleUpdateReward = async (formData: RewardFormData) => {
        if (!reward) return;
        
        try {
            setFormLoading(true);
            await rewardsService.updateReward(reward.id, formData);
            showSuccess('Prêmio atualizado com sucesso!');
            navigate('/MeusSorteios');
        } catch (err) {
            console.error('Erro ao atualizar prêmio:', err);
            
            // Extrai a mensagem de erro
            let errorMessage = 'Erro ao atualizar prêmio. Tente novamente.';
            if (err instanceof Error) {
                errorMessage = err.message;
            }
            
            showError(errorMessage);
        } finally {
            setFormLoading(false);
        }
    };

    const handleCancel = () => {
        navigate('/MeusSorteios');
    };

    if (loading) {
        return (
            <div className="reward-form-page-loading">
                <div className="loading-spinner"></div>
                <p>Carregando...</p>
            </div>
        );
    }

    return (
        <div className="reward-form-page">
            <div className="reward-form-page-container">
                <RewardForm
                    reward={reward}
                    onSubmit={reward ? handleUpdateReward : handleCreateReward}
                    onCancel={handleCancel}
                    loading={formLoading}
                />
            </div>
        </div>
    );
} 