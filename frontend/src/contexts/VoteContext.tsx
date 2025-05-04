import { createContext, useState, useContext } from "react";

interface VoteContextType {
  hasVoted: boolean;
  setHasVoted: (value: boolean) => void;
}

const VoteContext = createContext<VoteContextType | undefined>(undefined);

export const useVote = () => {
  const context = useContext(VoteContext);
  if (!context) {
    throw new Error("useVote must be used within a VoteProvider");
  }
  return context;
};

export default function VoteProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [hasVoted, setHasVoted] = useState(false);

  return (
    <VoteContext.Provider value={{ hasVoted, setHasVoted }}>
      {children}
    </VoteContext.Provider>
  );
}
