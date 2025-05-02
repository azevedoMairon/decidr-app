import { FormEvent } from "react";
import Participant from "../models/Participant";
import Button from "./Button";

interface Props {
  onSubmit?: (e: FormEvent<HTMLFormElement>) => void;
  participants: Participant[];
}

export default function VotingForm({ participants, onSubmit }: Props) {
  return (
    <form className="flex flex-col items-center" onSubmit={onSubmit}>
      {participants
        .filter((p) => p.isNominated)
        .map((p) => (
          <label className="w-[100%]">
            <input
              type="radio"
              name="participant"
              value={p.id}
              className="hidden peer"
            />
            <div className="flex items-center mb-6 justify-between border rounded-lg border-highlight px-28 transition ease-in-out cursor-pointer hover:bg-highlight hover:text-midnight peer-checked:bg-highlight peer-checked:text-dusk">
              <img
                src={p.imageUrl}
                alt={p.name}
                className="w-26 h-26 grayscale object-cover rounded-lg inline-block"
              />
              <span className="font-mono text-xl">{p.name}</span>
            </div>
          </label>
        ))}
      <Button text="Votar" type="submit" />
    </form>
  );
}
