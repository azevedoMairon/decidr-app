import { useParticipants } from "../contexts/ParticipantContext";
import AppLogo from "./AppLogo";
import ParticipantBubble from "./ParticipantBubble";

export default function Header() {
  const { participants, loading } = useParticipants();

  if (loading) return <></>

  return (
    <header className="flex flex-col items-center py-8">
      <AppLogo className="pb-8" />
      <span className="w-[100%] bg-highlight h-[1px]"></span>
      <div className="flex mt-6 justify-center">
        {participants.map((p) => (
          <ParticipantBubble key={p.id} participant={p} />
        ))}
      </div>
    </header>
  );
}
