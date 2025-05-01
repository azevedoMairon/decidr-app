import { useEffect, useState } from "react";
import AppLogo from "../components/AppLogo";
import Participant from "../models/Participant";

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
      <header className="absolute top-0 left-0 flex flex-col w-[100vw] items-center py-8">
        <AppLogo className="pb-8" />
        <span className="w-[100%] bg-highlight h-[1px]"></span>
        <div className="flex justify-center overflow-visible mt-6">
          <div className="flex relative">
            {participants.map((p) => (
              <div key={p.name} className="relative text-center">
                <img
                  src={p.imageUrl}
                  alt={p.name}
                  className="w-14 h-14 grayscale object-cover rounded-full inline-block hover:scale-125 transition-transform duration-100 origin-center cursor-pointer"
                />
              </div>
            ))}
          </div>
        </div>
      </header>
      <div className="grid grid-cols-2">
        <div className="flex justify-center translate-y-1/6">
          <p className="font-mono text-3xl text-justify text-light w-[65%] p-8">
            É você quem manda no jogo. Vote agora no participante que deve
            deixar o programa — o paredão está pegando fogo e cada clique conta.
          </p>
        </div>
        <div className="flex flex-col justify-items-stretch">
          {participants
            .filter((p) => p.isNominated)
            .map((p) => (
              <div className="flex items-center mb-6 justify-between w-[50%] border rounded-lg border-highlight px-16 transition ease-in-out cursor-pointer hover:bg-highlight hover:text-midnight">
                <img
                  src={p.imageUrl}
                  alt={p.name}
                  className="w-26 h-26 grayscale object-cover rounded-lg inline-block"
                />
                <span className="font-mono text-xl">{p.name}</span>
              </div>
            ))}
        </div>
      </div>
    </>
  );
}
