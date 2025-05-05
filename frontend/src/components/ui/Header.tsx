import { useParticipants } from "../../contexts/ParticipantContext";
import AppLogo from "./AppLogo";
import ParticipantBubble from "../participant/ParticipantBubble";
import CustomButton from "./CustomButton";
import { useLocation, useNavigate } from "react-router-dom";
import BackButton from "./BackButton";

export default function Header() {
  const { participants, loading } = useParticipants();
  const location = useLocation();
  const navigate = useNavigate();

  if (loading) return <></>;

  return (
    <header className="z-50 absolute top-0 left-0 flex flex-col items-center w-[100vw]">
      <div className="flex py-8 justify-between items-center w-[70%]">
        <BackButton />
        <AppLogo />
        {location.pathname != "/stats" ? (
          <CustomButton
            text="Ver Estatisticas"
            onClick={() => navigate("/stats")}
          />
        ) : (
          <div className="w-[156.83px]" />
        )}
      </div>
      <span className="w-[100%] bg-highlight h-[1px]"></span>
      <div className="flex mt-6 justify-center">
        {participants.map((p) => (
          <ParticipantBubble key={p.id} participant={p} />
        ))}
      </div>
    </header>
  );
}
