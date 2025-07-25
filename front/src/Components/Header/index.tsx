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
        {!isMobile && <a href="/MeuPerfil" className="nav-link cart-text">Meu Perfil</a>}
        <a href="/MeusSorteios" className="nav-link">Meus sorteios</a>
        <a href="/contato" className="nav-link">Contato</a>
      </div>
      <div style={{ display: 'flex', alignItems: 'center', paddingRight: '2rem' }}>
        <a href="/MeuPerfil" className="cart-icon cart-icon-mobile">
          <span className="material-icons">person</span>
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
