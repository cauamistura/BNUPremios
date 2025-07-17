import { useNavigate, useLocation } from 'react-router-dom';

export const useRedirectToLogin = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const redirectToLogin = () => {
    navigate('/auth/login', { state: { from: location } });
  };

  return { redirectToLogin };
}; 