import { createContext, useContext, useEffect, useState } from "react";
import Participant from "../models/Participant";

const PARTICIPANTS_STORAGE = "participants-cache";

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
    const cache = localStorage.getItem(PARTICIPANTS_STORAGE);

    if (cache) {
      const parsed = JSON.parse(cache);
      setParticipants(parsed.map((p: any) => new Participant(p)));
      setLoading(false);
      return;
    }

    fetch(`${import.meta.env.VITE_BACKEND_URL}/api/participants`)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch participants");
        return res.json();
      })
      .then((data) => {
        localStorage.setItem(PARTICIPANTS_STORAGE, JSON.stringify(data));
        setParticipants(data.map((p: any) => new Participant(p)));
        setLoading(false);
      });
  }, []);

  return (
    <ParticipantContext.Provider value={{ participants, loading }}>
      {children}
    </ParticipantContext.Provider>
  );
}
