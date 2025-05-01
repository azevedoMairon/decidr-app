import Participant from "../models/Participant";
import AppLogo from "./AppLogo";
import ParticipantBubble from "./ParticipantBubble";

interface Props {
    participants: Participant[];
}

export default function Header({ participants }: Props) {
  return (
    <header className="flex flex-col items-center py-8">
      <AppLogo className="pb-8" />
      <span className="w-[100%] bg-highlight h-[1px]"></span>
      <div className="flex mt-6 justify-center">
        {participants.map((p) => (
          <ParticipantBubble participant={p} />
        ))}
      </div>
    </header>
  );
}
