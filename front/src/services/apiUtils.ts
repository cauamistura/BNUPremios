// Utilitários para requisições API com autenticação

const API_BASE_URL = '/api/v1';

// Função para obter o token do localStorage
export const getAuthToken = (): string | null => {
  return localStorage.getItem('token');
};

// Função para criar headers com autenticação
export const getAuthHeaders = (): HeadersInit => {
  const token = getAuthToken();
  return {
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {})
  };
};

// Função para fazer logout quando token expirar
const handleTokenExpired = () => {
  // Limpar dados de autenticação
  localStorage.removeItem('user');
  localStorage.removeItem('token');
  
  // Redirecionar para login
  window.location.href = '/auth/login';
};

// Função para fazer requisições autenticadas
export const authenticatedFetch = async (
  endpoint: string,
  options: RequestInit = {}
): Promise<Response> => {
  const url = `${API_BASE_URL}${endpoint}`;
  const headers = {
    ...getAuthHeaders(),
    ...(options.headers || {})
  };

  const response = await fetch(url, {
    ...options,
    headers
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    const errorMessage = errorData.message || `Erro na requisição: ${response.status}`;
    
    // Verificar se é erro de token expirado
    if (errorMessage.includes('token is expired') || 
        errorMessage.includes('token has invalid claims') ||
        response.status === 401) {
      console.log('Token expirado detectado, fazendo logout...');
      handleTokenExpired();
    }
    
    throw new Error(errorMessage);
  }

  return response;
}; 