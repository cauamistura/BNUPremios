import React, { useState } from 'react';
import './index.css';

const Header: React.FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [cartCount] = useState(0);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const isMobile = window.innerWidth < 768;
  
  return (
    <header className="header-container">      
      <a href="/" className="logo">BNU Prêmios</a>        
      <div className="nav-menu" style={{ right: isMenuOpen ? '0' : '-100%' }}>             
        {!isMobile && <a href="/carrinho" className="nav-link cart-text">Carrinho</a>}                   
        <a href="/participar" className="nav-link">Participar</a>
        <a href="/contato" className="nav-link">Contato</a>
      </div>
      <div style={{ display: 'flex', alignItems: 'center', paddingRight: '2rem' }}>
        <a href="/carrinho" className="cart-icon cart-icon-mobile">
          <span className="material-icons">shopping_cart</span>
          {cartCount > 0 && <span className="cart-count">{cartCount}</span>}
        </a>
        <div className="menu-icon" onClick={toggleMenu}>
          <span className="material-icons">
            {isMenuOpen ? 'close' : 'menu'}
          </span>
        </div>
      </div>
    </header>
  );
};

export default Header;
