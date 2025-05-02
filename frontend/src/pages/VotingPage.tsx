import { useEffect, useState } from "react";
import Participant from "../models/Participant";
import Header from "../components/Header";
import VotingForm from "../components/VotingForm";

async function castVote(e: React.FormEvent<HTMLFormElement>) {
  e.preventDefault();

  const form = e.target as HTMLFormElement;
  const selected = new FormData(form).get("participant");

  console.log("Participante votado:", selected);

  try {
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/api/vote`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ participant: selected }),
      }
    );

    if (!response.ok) {
      throw new Error("Erro ao registrar o voto.");
    }

    const result = await response.json();

    console.log("Voto registrado com sucesso: ", result);
  } catch (err) {
    console.error(err);
  }
}

export default function VotePage() {
  const [participants, setParticipants] = useState<Participant[]>([]);

  useEffect(() => {
    fetch(`${import.meta.env.VITE_BACKEND_URL}/api/participants`)
      .then((res) => {
        if (!res.ok) throw new Error("Error while fetching participants.");
        return res.json();
      })
      .then((data) => {
        setParticipants(data.map((p: any) => new Participant(p)));
      });
  }, []);

  return (
    <>
      <Header participants={participants} />
      <div className="flex justify-center items-center h-[75%]">
        <div className="grid grid-cols-2 gap-32 w-[70%]">
          <div className="flex justify-center translate-y-1/6">
            <p className="font-mono text-3xl text-justify text-light p-8">
              É você quem manda no jogo. Vote agora no participante que deve
              deixar o programa — o paredão está pegando fogo e cada clique
              conta.
            </p>
          </div>
          <VotingForm participants={participants} onSubmit={castVote} />
        </div>
      </div>
    </>
  );
}
