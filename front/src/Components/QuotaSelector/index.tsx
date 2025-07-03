import React, { useState } from 'react';
import './index.css';

interface QuotaSelectorProps {
  price: number;
  minQuota: number;
  onChange?: (quantity: number) => void;
}

const options = [5, 10, 20, 50, 100];

const QuotaSelector: React.FC<QuotaSelectorProps> = ({ price, minQuota, onChange }) => {
  const [quantity, setQuantity] = useState(minQuota || 1);

  const handleSelect = (q: number) => {
    setQuantity(q);
    onChange?.(q);
  };

  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    let val = Number(e.target.value);
    if (isNaN(val) || val < 1) val = 1;
    setQuantity(val);
    onChange?.(val);
  };

  return (
    <div className="quota-selector">
      <div className="quota-options">
        {options.map(opt => (
          <button
            key={opt}
            type="button"
            className={`quota-option-btn${quantity === opt ? ' selected' : ''}`}
            onClick={() => handleSelect(opt)}
          >
            {opt}
          </button>
        ))}
        <input
          type="number"
          min={1}
          className="quota-option-input no-spinner"
          value={quantity}
          onChange={handleInput}
        />
      </div>
      <div className="quota-total">
        Total: <span>R$ {(quantity * price).toLocaleString('pt-BR', { minimumFractionDigits: 2 })}</span>
      </div>
    </div>
  );
};

export default QuotaSelector; 