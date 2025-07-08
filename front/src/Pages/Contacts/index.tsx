import React from 'react';
import './index.css';

const Contacts: React.FC = () => {
    return (
        <div className="contacts-container">

            <div className="contacts-content">
                <div className="profile-section">
                    <div className="profile-card">
                        <div className="profile-avatar">
                            <img 
                                src="https://avatars.githubusercontent.com/u/91917148?v=4" 
                                alt="Cauã Mistura" 
                            />
                        </div>
                        <div className="profile-info">
                            <h2>Cauã Mistura</h2>
                            <p className="profile-title">Desenvolvedor Full Stack</p>
                            <p className="profile-description">
                                Sou formado em Informática pelo IFSC e curso superior em analise e desenvolvimento de sistemas pela UNIVALI em quanto trabalho com programação desktop/web, especialmente utilizando c# e .Net, juntamente com Delphi.
                            </p>
                        </div>
                    </div>
                </div>

                <div className="contact-methods">
                    <div className="contact-card">
                        <div className="contact-icon">📧</div>
                        <div className="contact-info">
                            <h3>Email</h3>
                            <p>cauamistura1@gmail.com</p>
                            <a href="mailto:cauamistura1@gmail.com" className="contact-link">
                                Enviar Email
                            </a>
                        </div>
                    </div>

                    <div className="contact-card">
                        <div className="contact-icon">📱</div>
                        <div className="contact-info">
                            <h3>Telefone</h3>
                            <p>(47) 98921-4110</p>
                            <a href="tel:+5511999999999" className="contact-link">
                                Ligar Agora
                            </a>
                        </div>
                    </div>

                    <div className="contact-card">
                        <div className="contact-icon">💼</div>
                        <div className="contact-info">
                            <h3>LinkedIn</h3>
                            <p>caua-mistura</p>
                            <a 
                                href="https://www.linkedin.com/in/caua-mistura/" 
                                target="_blank" 
                                rel="noopener noreferrer"
                                className="contact-link"
                            >
                                Ver Perfil
                            </a>
                        </div>
                    </div>

                    <div className="contact-card">
                        <div className="contact-icon">🐙</div>
                        <div className="contact-info">
                            <h3>GitHub</h3>
                            <p>github.com/cauamistura</p>
                            <a 
                                href="https://github.com/cauamistura" 
                                target="_blank" 
                                rel="noopener noreferrer"
                                className="contact-link"
                            >
                                Ver Repositórios
                            </a>
                        </div>
                    </div>
                </div>              
            </div>
        </div>
    );
};

export default Contacts; 