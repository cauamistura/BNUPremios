import React from 'react';
import TopBuyerCard from './TopBuyerCard';
import type { Buyer } from '../../Models/Buyers';
import './index.css';

interface TopBuyersProps {
  buyers: Buyer[];
}

const TopBuyers: React.FC<TopBuyersProps> = ({ buyers }) => {
  if (!buyers || !Array.isArray(buyers) || buyers.length === 0) return null;

  // Ordenar por total_numbers (maior para menor) e pegar os top 3
  const sortedBuyers = [...buyers].sort((a, b) => b.total_numbers - a.total_numbers);
  const top3Buyers = sortedBuyers.slice(0, 3);
  const otherBuyers = sortedBuyers.slice(3);

  return (
    <div className="top-buyers-section">
      <div className="top-buyers-title">
        <span role="img" aria-label="trophy">🏆</span> Top compradores
      </div>
      <div className="top-buyers-subtitle">TOP COMPRADOR GERAL</div>
      
      {/* Top 3 com cards especiais */}
      <div className="top-buyers-list-container">
        <div className="top-buyers-list">
          {top3Buyers.map((buyer, idx) => (
            buyer && buyer.user && typeof buyer.total_numbers === 'number' ? (
              <TopBuyerCard
                key={buyer.user.id}
                name={buyer.user.name}
                amount={buyer.total_numbers}
                position={idx + 1}
              />
            ) : null
          ))}
        </div>
      </div>

      {/* Lista simples dos demais compradores */}
      {otherBuyers.length > 0 && (
        <div className="other-buyers-section">
          <div className="other-buyers-title">Outros Participantes</div>
          <div className="other-buyers-list">
            {otherBuyers.map((buyer, idx) => (
              buyer && buyer.user && typeof buyer.total_numbers === 'number' ? (
                <div key={buyer.user.id} className="other-buyer-item">
                  <span className="other-buyer-position">{idx + 4}º</span>
                  <span className="other-buyer-name">{buyer.user.name}</span>
                  <span className="other-buyer-numbers">{buyer.total_numbers} números</span>
                </div>
              ) : null
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default TopBuyers; 