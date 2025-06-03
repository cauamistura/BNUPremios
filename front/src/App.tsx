import Header from "./Components/Header";
import "./App.css";
import { BrowserRouter as Router } from "react-router-dom";
import AppRoutes from "./routes";

function App() {
  return (
    <Router>
      <div className="app-container">
        <Header />
        <div className="content-container">
          <AppRoutes />
        </div>
      </div>
    </Router>
  );
}

export default App;