import React from 'react';
import './index.css';

const Header: React.FC = () => {
  return (
    <header className="header-container">      
      <a href="/" className="logo">BNU Prêmios</a>        
      <nav className="nav-menu">        
        <a href="/premios" className="nav-link">Prêmios</a>
        <a href="/participar" className="nav-link">Participar</a>
        <a href="/contato" className="nav-link">Contato</a>
      </nav>
    </header>
  );
};

export default Header;
