import { useParticipants } from "../contexts/ParticipantContext";
import { useVote } from "../contexts/VoteContext";
import toast from "react-hot-toast";
import Button from "./ui/Button";
import { postVote } from "../services/api";
import ParticipantCard from "./participant/ParticipantCard";

export default function VoteForm() {
  const { participants, loading } = useParticipants();
  const { setHasVoted } = useVote();

  const castVote = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const selectedParticipantId = formData.get("participant");

    if (!selectedParticipantId) {
      toast.error("Selecione um participante antes de votar.");
      return;
    }

    try {
      await postVote(selectedParticipantId);
      setHasVoted(true);
      toast.success("Voto registrado com sucesso!");
    } catch (error) {
      toast.error("Falha ao enviar o voto.");
    }
  };

  if (loading) return <></>;

  const nominated = participants.filter((p) => p.isNominated);

  return (
    <form className="flex flex-col items-center" onSubmit={castVote}>
      {nominated.map((p) => (
        <label key={p.id} className="w-[100%]">
          <input
            type="radio"
            name="participant"
            value={p.id}
            className="hidden peer"
          />
          <ParticipantCard
            name={p.name}
            imageUrl={p.imageUrl}
            className="ease-in-out cursor-pointer hover:bg-highlight hover:text-midnight peer-checked:bg-highlight peer-checked:text-dusk"
          />
        </label>
      ))}
      <Button text="Votar" type="submit" />
    </form>
  );
}
