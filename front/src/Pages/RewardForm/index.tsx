import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import RewardForm from "../../Components/RewardForm";
import type { RewardFormData } from "../../Components/RewardForm";
import { rewardsService } from "../../services/rewardsService";
import type { Reward } from "../../Models/Reaward";
import { useAuth } from "../../hooks/useAuth";
import "./index.css";

export default function RewardFormPage() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const { user } = useAuth();
    
    const [reward, setReward] = useState<Reward | undefined>();
    const [loading, setLoading] = useState(true);
    const [formLoading, setFormLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchReward = async () => {
            if (!id) {
                setLoading(false);
                return;
            }

            try {
                setLoading(true);
                setError(null);
                const data = await rewardsService.getRewardById(id);
                setReward(data);
            } catch (err) {
                setError('Erro ao carregar o prêmio. Tente novamente.');
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
            navigate('/MeusSorteios');
        } catch (err) {
            console.error('Erro ao criar prêmio:', err);
            setError('Erro ao criar prêmio. Tente novamente.');
        } finally {
            setFormLoading(false);
        }
    };

    const handleUpdateReward = async (formData: RewardFormData) => {
        if (!reward) return;
        
        try {
            setFormLoading(true);
            await rewardsService.updateReward(reward.id, formData);
            navigate('/MeusSorteios');
        } catch (err) {
            console.error('Erro ao atualizar prêmio:', err);
            setError('Erro ao atualizar prêmio. Tente novamente.');
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

    if (error) {
        return (
            <div className="reward-form-page-error">
                <p className="error-message">{error}</p>
                <button className="error-reload-btn" onClick={() => window.location.reload()}>
                    Tentar novamente
                </button>
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