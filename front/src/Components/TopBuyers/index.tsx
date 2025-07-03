import React from 'react';
import TopBuyerCard from './TopBuyerCard';
import './index.css';

interface Buyer {
  name: string;
  amount: number;
}

interface TopBuyersProps {
  buyers: Buyer[];
}

const TopBuyers: React.FC<TopBuyersProps> = ({ buyers }) => {
  if (!buyers || buyers.length === 0) return null;
  return (
    <div className="top-buyers-section">
      <div className="top-buyers-title">
        <span role="img" aria-label="trophy">🏆</span> Top compradores
      </div>
      <div className="top-buyers-subtitle">TOP COMPRADOR GERAL</div>
      <div className="top-buyers-list-container">
        <div className="top-buyers-list">
          {buyers.slice(0, 3).map((buyer, idx) => (
            <TopBuyerCard
              key={buyer.name}
              name={buyer.name}
              amount={buyer.amount}
              position={idx + 1}
            />
          ))}
        </div>
      </div>

    </div>
  );
};

export default TopBuyers; 