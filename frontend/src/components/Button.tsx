interface Props {
  text: string;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
}

export default function Button({ text, onClick, type = "button"}: Props) {
  return (
    <>
      <button
        onClick={onClick}
        className="mt-6 px-4 py-2 border-2 w-[50%] border-highlight light font-bold hover:bg-highlight hover:text-midnight"
        type={type}
      >
        {text}
      </button>
    </>
  );
}
