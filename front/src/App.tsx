import Header from "./Components/Header";
import "./App.css";
import { BrowserRouter as Router } from "react-router-dom";
import AppRoutes from "./routes";
import { AuthProvider } from "./contexts/AuthContext";
import { ToastProvider } from "./contexts/ToastContext";
import ToastContainer from "./Components/ToastContainer";
import { useToastContext } from "./contexts/ToastContext";

function AppContent() {
  const { toasts, removeToast } = useToastContext();

  return (
    <div className="app-container">
      <Header />
      <div className="content-container">
        <AppRoutes />
      </div>
      <ToastContainer toasts={toasts} onRemove={removeToast} />
    </div>
  );
}

function App() {
  return (
    <AuthProvider>
      <ToastProvider>
        <Router>
          <AppContent />
        </Router>
      </ToastProvider>
    </AuthProvider>
  );
}

export default App;