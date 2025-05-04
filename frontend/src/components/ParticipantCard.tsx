interface ParticipantCardProps {
  name: string;
  imageUrl: string;
  percent?: number;
  className?: string;
}

export default function ParticipantCard({
  name,
  imageUrl,
  percent,
  className,
}: ParticipantCardProps) {
  return (
    <div
      className={`flex items-center mb-6 justify-between border rounded-lg border-highlight px-28 transition w-full ${className}`}
      style={{
        background: `linear-gradient(to right, rgb(45 71 84) ${percent}%, rgb(9 26 35) ${percent}%)`,
      }}
    >
      <img
        src={imageUrl}
        alt={name}
        className="w-26 h-26 grayscale object-cover rounded-lg inline-block"
      />
      <span className="font-mono text-xl">
        {name}
        {percent ? <>- {percent}%</> : <></>}
      </span>
    </div>
  );
}
