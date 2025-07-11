import React, { useState } from 'react';
import { useAuth } from '../../../hooks/useAuth';
import { useNavigate } from 'react-router-dom';
import './index.css';

const LoginPage: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const { login } = useAuth();
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');
        setIsLoading(true);
        try {
            const success = await login(email, password);
            if (success) {
                navigate('/profile'); // Redireciona para a página de profile
            } else {
                setError('Email ou senha incorretos');
            }
        } catch {
            setError('Erro ao fazer login. Tente novamente.');
        } finally {
            setIsLoading(false);
        }
    };

    const handleRegisterClick = () => {
        navigate('/auth/register');
    };

    return (
        <div className="auth-page">
            <div className="auth-container">
                <div className="auth-header">
                    <h1>Login</h1>
                    <p>Entre com suas credenciais</p>
                </div>

                <form onSubmit={handleSubmit} className="auth-form">
                    <div className="form-group">
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
                    <div className="form-group">
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
                        <div className="auth-error">
                            {error}
                        </div>
                    )}
                    <button 
                        type="submit" 
                        className="auth-submit-btn"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Entrando...' : 'Entrar'}
                    </button>
                </form>

                <div className="auth-footer">
                    <p>Não tem uma conta?</p>
                    <button 
                        className="auth-link-btn" 
                        onClick={handleRegisterClick}
                        disabled={isLoading}
                    >
                        Criar conta
                    </button>
                </div>
            </div>
        </div>
    );
};

export default LoginPage; 