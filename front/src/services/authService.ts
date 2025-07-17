// Serviço de autenticação para login e registro

const API_URL = '/api/v1/auth';

export interface RegisterPayload {
  name: string;
  email: string;
  password: string;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export async function login({ email, password }: LoginPayload) {
  const response = await fetch(`${API_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  if (!response.ok) {
    throw new Error('Credenciais inválidas');
  }
  return await response.json(); // { token, user }
}

export async function register({ name, email, password }: RegisterPayload) {
  console.log(name, email, password);

  const response = await fetch(`${API_URL}/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, email, password })
  });
  if (!response.ok) {
    throw new Error('Erro ao registrar usuário');
  }
  return await response.json(); // { id, name, email, ... }
} 