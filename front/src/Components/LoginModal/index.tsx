import React, { useState } from 'react';
import { useAuth } from '../../hooks/useAuth';
import './index.css';

interface LoginModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSuccess: () => void;
}

const LoginModal: React.FC<LoginModalProps> = ({ isOpen, onClose, onSuccess }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const { login } = useAuth();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setIsLoading(true);

        try {
            const success = await login(email, password);
            if (success) {
                onSuccess();
                onClose();
                setEmail('');
                setPassword('');
            } else {
                setError('Email ou senha incorretos');
            }
        } catch {
            setError('Erro ao fazer login. Tente novamente.');
        } finally {
            setIsLoading(false);
        }
    };

    const handleClose = () => {
        if (!isLoading) {
            setEmail('');
            setPassword('');
            setError('');
            onClose();
        }
    };

    if (!isOpen) return null;

    return (
        <div className="login-modal-overlay" onClick={handleClose}>
            <div className="login-modal" onClick={(e) => e.stopPropagation()}>
                <div className="login-modal-header">
                    <h2>Login</h2>
                    <button 
                        className="login-modal-close" 
                        onClick={handleClose}
                        disabled={isLoading}
                    >
                        ×
                    </button>
                </div>

                <form onSubmit={handleSubmit} className="login-form">
                    <div className="login-form-group">
                        <label htmlFor="email">Email</label>
                        <input
                            type="email"
                            id="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="Digite seu email"
                            required
                            disabled={isLoading}
                        />
                    </div>

                    <div className="login-form-group">
                        <label htmlFor="password">Senha</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="Digite sua senha"
                            required
                            disabled={isLoading}
                        />
                    </div>

                    {error && (
                        <div className="login-error">
                            {error}
                        </div>
                    )}

                    <button 
                        type="submit" 
                        className="login-submit-btn"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Entrando...' : 'Entrar'}
                    </button>
                </form>

                <div className="login-modal-footer">
                    <p>Credenciais de teste:</p>
                    <p><strong>Email:</strong> joao.silva@email.com</p>
                    <p><strong>Senha:</strong> 123456</p>
                </div>
            </div>
        </div>
    );
};

export default LoginModal; 