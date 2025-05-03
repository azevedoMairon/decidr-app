import { useParticipants } from "../contexts/ParticipantContext";
import toast from "react-hot-toast";
import Button from "./Button";

async function castVote(e: React.FormEvent<HTMLFormElement>) {
  e.preventDefault();

  const form = e.target as HTMLFormElement;
  const formData = new FormData(form);
  const selectedParticipantId = formData.get("participant");

  if (!selectedParticipantId) {
    toast.error("Selecione um participante antes de votar.");
    return;
  }

  try {
    const payload = { participant_id: selectedParticipantId };
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/api/vote`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      }
    );

    if (!response.ok) throw new Error("Erro ao registrar o voto.");
    toast.success("Voto registrado com sucesso!");

  } catch (error) {

    toast.error("Falha ao enviar o voto.");
    console.error("Falha ao enviar voto:", error);
  }
}

export default function VotingForm() {
  const { participants, loading } = useParticipants();

  if (loading) return <></>;

  return (
    <form className="flex flex-col items-center" onSubmit={castVote}>
      {participants
        .filter((p) => p.isNominated)
        .map((p) => (
          <label className="w-[100%]">
            <input
              type="radio"
              name="participant"
              value={p.id}
              className="hidden peer"
            />
            <div className="flex items-center mb-6 justify-between border rounded-lg border-highlight px-28 transition ease-in-out cursor-pointer hover:bg-highlight hover:text-midnight peer-checked:bg-highlight peer-checked:text-dusk">
              <img
                src={p.imageUrl}
                alt={p.name}
                className="w-26 h-26 grayscale object-cover rounded-lg inline-block"
              />
              <span className="font-mono text-xl">{p.name}</span>
            </div>
          </label>
        ))}
      <Button text="Votar" type="submit" />
    </form>
  );
}
