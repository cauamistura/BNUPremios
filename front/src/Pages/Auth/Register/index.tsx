import React, { useState } from 'react';
import { useAuth } from '../../../hooks/useAuth';
import { useNavigate } from 'react-router-dom';
import './index.css';

const RegisterPage: React.FC = () => {
  const { register } = useAuth();
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);
    
    try {
      const ok = await register({ name, email, password });
      
      if (ok) {
        // Redireciona para a página de login após registro bem-sucedido
        navigate('/auth/login');
      } else {
        setError('Erro ao registrar. Tente novamente.');
      }
    } catch {
      setError('Erro ao registrar. Tente novamente.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleLoginClick = () => {
    navigate('/auth/login');
  };

  return (
    <div className="auth-page">
      <div className="auth-container">
        <div className="auth-header">
          <h1>Criar Conta</h1>
          <p>Preencha os dados para se cadastrar</p>
        </div>

        <form className="auth-form">
          <div className="form-group">
            <label htmlFor="register-name">Nome</label>
            <input 
              id="register-name" 
              value={name} 
              onChange={e => setName(e.target.value)} 
              placeholder="Digite seu nome completo"
              required 
              disabled={isLoading} 
            />
          </div>
          <div className="form-group">
            <label htmlFor="register-email">Email</label>
            <input 
              id="register-email" 
              type="email" 
              value={email} 
              onChange={e => setEmail(e.target.value)} 
              placeholder="Digite seu email"
              required 
              disabled={isLoading} 
            />
          </div>
          <div className="form-group">
            <label htmlFor="register-password">Senha</label>
            <input 
              id="register-password" 
              type="password" 
              value={password} 
              onChange={e => setPassword(e.target.value)} 
              placeholder="Digite sua senha"
              required 
              disabled={isLoading} 
            />
          </div>
          {error && <div className="auth-error">{error}</div>}
          <button 
            onClick={handleSubmit}  
            className="auth-submit-btn" 
            disabled={isLoading}
          >
            {isLoading ? 'Criando conta...' : 'Criar conta'}
          </button>
        </form>

        <div className="auth-footer">
          <p>Já tem uma conta?</p>
          <button 
            className="auth-link-btn" 
            onClick={handleLoginClick}
            disabled={isLoading}            
          >
            Fazer login
          </button>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage; 