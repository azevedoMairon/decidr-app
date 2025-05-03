import Header from "../components/Header";
import VotingForm from "../components/VotingForm";

export default function VotePage() {
  return (
    <>
      <Header />
      <div className="flex justify-center items-center h-[75%]">
        <div className="grid grid-cols-2 gap-32 w-[70%]">
          <div className="flex justify-center translate-y-1/6">
            <p className="font-mono text-3xl text-justify text-light p-8">
              É você quem manda no jogo. Vote agora no participante que deve
              deixar o programa — o paredão está pegando fogo e cada clique
              conta.
            </p>
          </div>
          <VotingForm />
        </div>
      </div>
    </>
  );
}
