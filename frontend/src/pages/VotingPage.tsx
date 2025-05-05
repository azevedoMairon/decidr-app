import Header from "../components/ui/Header";
import VoteResults from "../components/VoteResults";
import VoteForm from "../components/VoteForm";
import { useVote } from "../contexts/VoteContext";

export default function VotePage() {
  const { hasVoted } = useVote();

  return (
    <>
      <Header />
      <main className="z-0 absolute top-0 left-0 w-full h-full flex justify-center items-center">
        <div className="grid grid-cols-2 gap-32 w-[70%] pt-[260px]">
          <div className="flex justify-center">
            <p className="font-mono text-3xl text-justify text-light p-8">
              É você quem manda no jogo. Vote agora no participante que deve
              deixar o programa — o paredão está pegando fogo e cada clique
              conta.
            </p>
          </div>
          {!hasVoted ? <VoteForm /> : <VoteResults />}
        </div>
      </main>
    </>
  );
}
