import Header from "./Components/Header";
import "./App.css";
import RewardCardList from "./Components/RewardCardList";
import rewardsMock from "./assets/Mocks/Rewards.json";

function App() {
  return (
    <div className="app-container">
      <Header />
      <div className="content-container">
        <RewardCardList rewards={rewardsMock} />
      </div>
    </div>
  );
}

export default App;