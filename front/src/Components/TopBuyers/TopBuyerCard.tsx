import React from 'react';

interface TopBuyerCardProps {
  name: string;
  amount: number;
  position: number;
}

const medalEmoji = ["🥇", "🥈", "🥉"];

const TopBuyerCard: React.FC<TopBuyerCardProps> = ({ name, amount, position }) => {
  return (
    <div className={`top-buyer-card position-${position}`}>
      <div className="top-buyer-rank">
        {position}º{' '}
        <span role="img" aria-label="medal">
          {medalEmoji[position - 1] || ''}
        </span>
      </div>
      <div className="top-buyer-name">{name.toUpperCase()}</div>
      <div className="top-buyer-amount">{amount} COTAS</div>
    </div>
  );
};

export default TopBuyerCard; 