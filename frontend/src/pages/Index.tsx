import { useNavigate } from "react-router-dom";
import AppLogo from "../components/ui/AppLogo";
import CustomButton from "../components/ui/CustomButton";

export default function Index() {
  const navigate = useNavigate();

  return (
    <>
      <div className="flex flex-col items-center justify-between w-[100vw] h-[18vh]">
        <div className="text-center">
          <AppLogo />
          <p className="text-dusk-90 mt-2 font-mono">
            seu app de decisões em tempo real.
          </p>
        </div>

        <div className="w-[25%] flex justify-center">
          <CustomButton
            text="Votar Agora"
            onClick={() => navigate("/voting-page")}
          />
        </div>
      </div>
    </>
  );
}
