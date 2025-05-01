interface Props {
  text: string;
  onClick?: () => void;
}

export default function Button({ text, onClick }: Props) {
  return (
    <>
      <button
        onClick={onClick}
        className="mt-6 px-4 py-2 border-2 w-[15%] border-highlight light font-bold hover:bg-highlight hover:text-midnight"
      >
        {text}
      </button>
    </>
  );
}
