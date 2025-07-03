import { Routes, Route } from 'react-router-dom';
import Home from '../Pages/Home';
import Contacts from '../Pages/Contacts';
import RewardDetails from '../Pages/RewardDetails';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/contato" element={<Contacts />} />
            <Route path="/premio/:id" element={<RewardDetails />} />
            <Route path="/participar" element={<div>Página em construção</div>} />
            <Route path="/carrinho" element={<div>Página em construção</div>} />
        </Routes>
    );
};

export default AppRoutes; 