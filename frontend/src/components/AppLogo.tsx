import { motion } from "framer-motion";

interface Props {
    className?: string;
}
  

export default function AppLogo({ className }: Props) {
  return (
    <>
      <motion.div
        layoutId="appLogo"
        transition={{ type: "spring", stiffness: 70, damping: 15 }}
        className={className}
      >
        <h1 className="font-display text-6xl font-bold text-light">
          DECID<span className="text-highlight">R</span>.
        </h1>
      </motion.div>
    </>
  );
}
