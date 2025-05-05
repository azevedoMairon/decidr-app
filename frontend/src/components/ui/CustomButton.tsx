interface Props {
  text: string;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
}

export default function CustomButton({ text, onClick, type = "button" }: Props) {
  return (
    <button
      onClick={onClick}
      className="px-4 py-2 border-2 border-highlight font-bold hover:bg-highlight hover:text-midnight focus:outline-none"
      type={type}
    >
      {text}
    </button>
  );
}
