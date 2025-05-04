import Participant from "../models/Participant";

interface Props {
  participant: Participant;
}

export default function ParticipantBubble({ participant }: Props) {
  return (
    <img
      src={participant.imageUrl}
      alt={participant.name}
      className="w-14 h-14 grayscale object-cover rounded-full cursor-pointer transition hover:scale-125"
    />
  );
}
