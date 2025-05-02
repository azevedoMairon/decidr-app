import { motion } from "framer-motion";

interface Props {
  text: string;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
}

export default function Button({ text, onClick, type = "button" }: Props) {
  return (
    <motion.div
      layoutId="appButton"
      transition={{ type: "spring", stiffness: 70, damping: 15 }}
    >
      <button
        onClick={onClick}
        className="mt-6 px-4 py-2 border-2 border-highlight light font-bold hover:bg-highlight hover:text-midnight"
        type={type}
      >
        {text}
      </button>
    </motion.div>
  );
}
