import { useEffect, useState } from "react";
import { format } from "date-fns";
import Header from "../components/ui/Header";
import ParticipantCard from "../components/participant/ParticipantCard";
import VotesPerHourChart from "../components/VotePerHourChart";

interface VoteEntry {
  participant_id: string;
  count: number;
  hour: string;
}

interface Participant {
  id: string;
  name: string;
  image_url: string;
}

export default function VoteStats() {
  const [data, setData] = useState<VoteEntry[]>([]);
  const [participants, setParticipants] = useState<Participant[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchData() {
      const [resVotes, resParticipants] = await Promise.all([
        fetch(`${import.meta.env.VITE_BACKEND_URL}/api/results?byHour=true`),
        fetch(
          `${
            import.meta.env.VITE_BACKEND_URL
          }/api/participants?isNominated=true`
        ),
      ]);
      const votes: VoteEntry[] = await resVotes.json();
      const parts: Participant[] = await resParticipants.json();
      setData(votes);
      setParticipants(parts);
      setLoading(false);
    }
    fetchData();
  }, []);

  if (loading) return <p>Carregando estatísticas...</p>;

  const totalVotes = data.reduce((acc, d) => acc + d.count, 0);

  const votesByParticipant = data.reduce<Record<string, number>>((acc, d) => {
    acc[d.participant_id] = (acc[d.participant_id] || 0) + d.count;
    return acc;
  }, {});

  const votesByHour = data.reduce<Record<string, number>>((acc, d) => {
    const hour = format(new Date(d.hour), "HH:mm");
    acc[hour] = (acc[hour] || 0) + d.count;
    return acc;
  }, {});

  return (
    <>
      <Header />
      <main className="absolute inset-0 flex justify-center items-start overflow-hidden">
        <div className="flex flex-col items-center justify-center w-[100%] h-[100%]">
          <div className="p-6 space-y-8 text-white w-[70%] max-h-[60%] overflow-y-auto mt-28 flex flex-col items-center">
            <div className="text-2xl">
              <strong>Total Geral de Votos:</strong> {totalVotes}
            </div>
            <h2 className="text-2xl font-semibold">Votos por Participante</h2>
            <div className="space-y-2 w-[55%]">
              {participants.map((p) => (
                <ParticipantCard
                  key={p.id}
                  name={p.name}
                  imageUrl={p.image_url}
                  votes={votesByParticipant[p.id]}
                />
              ))}
            </div>

            <VotesPerHourChart votesByHour={votesByHour} />
          </div>
        </div>
      </main>
    </>
  );
}
