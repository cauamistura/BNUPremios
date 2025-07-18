import { useState, useEffect } from 'react';
import RewardCardList from "../../Components/RewardCardList";
import { rewardsService } from "../../services/rewardsService";
import type { Reward } from "../../Models/Reaward";
import { useToastContext } from "../../contexts/ToastContext";
import "./index.css";

export default function Home() {
    const { showError } = useToastContext();
    const [rewards, setRewards] = useState<Reward[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchRewards = async () => {
            try {
                setLoading(true);
                const response = await rewardsService.getRewards();
                // Garantir que rewards seja sempre um array válido
                setRewards(response.rewards || []);
            } catch (err) {
                const errorMessage = 'Erro ao carregar os prêmios. Tente novamente.';
                showError(errorMessage);
                console.error('Erro ao buscar prêmios:', err);
                // Em caso de erro, garantir que rewards seja um array vazio
                setRewards([]);
            } finally {
                setLoading(false);
            }
        };

        fetchRewards();
    }, []);

    if (loading) {
        return (
            <div className="home-loading-container">
                <div className="home-loading-spinner"></div>
                <p>Carregando prêmios...</p>
            </div>
        );
    }

    // Verificação adicional de segurança
    const rewardsArray = Array.isArray(rewards) ? rewards : [];

    return (
        <div className="home-container">             
            {rewardsArray.length > 0 ? (
                <RewardCardList rewards={rewardsArray} routeItem={"premio"} />
            ) : (
                <div className="home-empty-container">
                    <p>Não há prêmios disponíveis.</p>
                </div>
            )}
        </div>
    );
}