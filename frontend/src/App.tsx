import { useNavigate } from "react-router-dom";
import AppLogo from "./components/ui/AppLogo";
import Button from "./components/ui/Button";

export default function App() {
  const navigate = useNavigate();

  return (
    <>
      <div className="flex flex-col items-center justify-center w-[100vw] h-[100vh]">
        <AppLogo />

        <p className="text-dusk-90 mt-4 font-mono">
          seu app de decisões em tempo real.
        </p>

        <div className="w-[25%] flex justify-center">
          <Button text="Votar Agora" onClick={() => navigate("/voting-page")} />
        </div>
      </div>
    </>
  );
}
