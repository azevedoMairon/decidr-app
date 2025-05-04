import { useEffect, useState } from "react";
import { useParticipants } from "../contexts/ParticipantContext";
import { useVote } from "../contexts/VoteContext";
import { useNavigate } from "react-router-dom";
import Button from "./Button";
import { getResults } from "../services/api";
import toast from "react-hot-toast";
import ParticipantCard from "./ParticipantCard";

export default function VoteResults() {
  const { hasVoted, setHasVoted } = useVote();
  const navigate = useNavigate();
  const { participants, loading } = useParticipants();
  const [results, setResults] = useState<Record<string, number>>({});

  useEffect(() => {
    getResults()
      .then((data) => {
        setResults(data);
      })
      .catch(() => {
        toast.error("Erro ao buscar resultados");
      });
  }, []);

  const handleVoteAgain = () => {
    setHasVoted(false);
    console.log(hasVoted)
    navigate("/voting-page");
  };

  if (loading) return <></>;

  const nominated = participants.filter((p) => p.isNominated);

  return (
    <>
      <div className="flex flex-col items-center w-full">
        {nominated.map((p) => {
          const percent = results[p.id] ?? 0;
          return (
            <ParticipantCard
              key={p.id}
              name={p.name}
              imageUrl={p.imageUrl}
              percent={percent}
            />
          );
        })}
        <Button text="Votar Novamente" onClick={handleVoteAgain} />
      </div>
    </>
  );
}
