import { createContext, useContext, useEffect, useState } from "react";
import Participant from "../models/Participant";
import { getParticipants } from "../services/api";
import toast from "react-hot-toast";

interface ParticipantContextType {
  participants: Participant[];
  loading: boolean;
}

const ParticipantContext = createContext<ParticipantContextType>({
  participants: [],
  loading: true,
});

export const useParticipants = () => useContext(ParticipantContext);

export default function ParticipantProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [participants, setParticipants] = useState<Participant[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getParticipants()
      .then((data) => {
        setParticipants(data.map((p: any) => new Participant(p)));
      })
      .catch(() => {
        toast.error("Erro ao buscar participantes");
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  return (
    <ParticipantContext.Provider value={{ participants, loading }}>
      {children}
    </ParticipantContext.Provider>
  );
}
