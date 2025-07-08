import { Routes, Route } from 'react-router-dom';
import Home from '../Pages/Home';
import Contacts from '../Pages/Contacts';
import RewardDetails from '../Pages/RewardDetails';
import Profile from '../Pages/Profile';
import ProtectedRoute from '../Components/ProtectedRoute';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/contato" element={<Contacts />} />
            <Route path="/premio/:id" element={<RewardDetails />} />
            <Route path="/participar" element={<div>Página em construção</div>} />
            <Route path="/MeuPerfil" element={
                <ProtectedRoute>
                    <Profile />
                </ProtectedRoute>
            } />
        </Routes>
    );
};

export default AppRoutes; 