import { useState, useEffect } from 'react';
import RewardCardList from "../../Components/RewardCardList";
import { rewardsService } from "../../services/rewardsService";
import type { Reward } from "../../Models/Reaward";
import "./index.css";

export default function Home() {
    const [rewards, setRewards] = useState<Reward[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchRewards = async () => {
            try {
                setLoading(true);
                setError(null);
                const response = await rewardsService.getRewards();
                setRewards(response.rewards);
            } catch (err) {
                setError('Erro ao carregar os prêmios. Tente novamente.');
                console.error('Erro ao buscar prêmios:', err);
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

    if (error) {
        return (
            <div className="home-error-container">
                <p className="home-error-message">{error}</p>
                <button className="home-error-reload-btn" onClick={() => window.location.reload()}>
                    Tentar novamente
                </button>
            </div>
        );
    }

    return (
        <div className="home-container">            
            <RewardCardList rewards={rewards} />
        </div>
    );
}