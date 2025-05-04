import { useEffect, useState } from "react";
import { useParticipants } from "../contexts/ParticipantContext";
import { useNavigate } from "react-router-dom";
import Button from "./Button";

interface Props {
  setArg: (set: boolean) => void;
}

export default function VoteResults({ setArg }: Props) {
  const navigate = useNavigate();
  const { participants, loading } = useParticipants();
  const [results, setResults] = useState<Record<string, number>>({});

  useEffect(() => {
    const fetchResults = async () => {
      const res = await fetch(
        `${import.meta.env.VITE_BACKEND_URL}/api/results`
      );
      const data = await res.json();

      const totalVotes = data.reduce(
        (acc: number, curr: any) => acc + curr.count,
        0
      );
      const formatted = Object.fromEntries(
        data.map((r: any) => [
          r.participant_id,
          Math.round((r.count / totalVotes) * 100),
        ])
      );

      setResults(formatted);
    };

    fetchResults();
  }, []);

  if (loading) return <></>;
  const nominated = participants.filter((p) => p.isNominated);

  return (
    <>
      <div className="flex flex-col items-center w-full">
        {nominated.map((p) => {
          const percent = results[p.id] ?? 0;
          return (
            <>
              <div
                key={p.id}
                className="flex items-center mb-6 justify-between border rounded-lg border-highlight px-28 transition w-full text-white"
                style={{
                  background: `linear-gradient(to right, rgb(45 71 84) ${percent}%, rgb(9 26 35) ${percent}%)`,
                }}
              >
                <img
                  src={p.imageUrl}
                  alt={p.name}
                  className="w-26 h-26 grayscale object-cover rounded-lg inline-block"
                />
                <span className="font-mono text-xl text-shadow-lg">
                  {p.name} - {percent}%
                </span>
              </div>
            </>
          );
        })}
        <Button
          text="Votar Novamente"
          onClick={() => {
            navigate("/voting-page"), setArg(false);
          }}
        />
      </div>
    </>
  );
}
