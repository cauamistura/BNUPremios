import { Routes, Route } from 'react-router-dom';
import Home from '../Pages/Home';
import Contacts from '../Pages/Contacts';
import RewardDetails from '../Pages/RewardDetails';
import Profile from '../Pages/Profile';
import MyRewards from '../Pages/MyRewards';
import RewardFormPage from '../Pages/RewardForm';
import ProtectedRoute from '../Components/ProtectedRoute';
import RegisterPage from '../Pages/Auth/Register';
import LoginPage from '../Pages/Auth/Login';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/contato" element={<Contacts />} />
            <Route path="/premio/:id" element={<RewardDetails />} />
            <Route path="/MeusSorteios" element={
                <ProtectedRoute>
                    <MyRewards />
                </ProtectedRoute>
            } />
            <Route path="/MeuSorteio/novo" element={
                <ProtectedRoute>
                    <RewardFormPage />
                </ProtectedRoute>
            } />
            <Route path="/MeuSorteio/:id" element={
                <ProtectedRoute>
                    <RewardFormPage />
                </ProtectedRoute>
            } />
            <Route path="/auth/login" element={<LoginPage />} />
            <Route path="/auth/register" element={<RegisterPage />} />
            <Route path="/MeuPerfil" element={
                <ProtectedRoute>
                    <Profile />
                </ProtectedRoute>
            } />
        </Routes>
    );
};

export default AppRoutes; 