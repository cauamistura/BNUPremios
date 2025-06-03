import RewardCardList from "../../Components/RewardCardList";
import rewardsMock from "../../assets/Mocks/Rewards.json";

export default function Home() {
    return (
        <RewardCardList rewards={rewardsMock} />  
    )
}