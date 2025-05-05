import { Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  Legend,
} from "chart.js";

ChartJS.register(BarElement, CategoryScale, LinearScale, Tooltip, Legend);

export default function VotesPerHourChart({
  votesByHour,
}: {
  votesByHour: Record<string, number>;
}) {
  const sortedEntries = Object.entries(votesByHour).sort(([a], [b]) =>
    a.localeCompare(b)
  );

  const labels = sortedEntries.map(([hour]) => `${hour}h`);
  const dataValues = sortedEntries.map(([, count]) => count);

  const data = {
    labels,
    datasets: [
      {
        label: "Votos por hora",
        data: dataValues,
        backgroundColor: "rgb(45 71 84)",
        borderRadius: 4,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: { display: false },
    },
    scales: {
      y: {
        beginAtZero: true,
        ticks: {
          precision: 0,
        },
      },
    },
  };

  return (
    <div className="w-full max-w-3xl mx-auto mt-6">
      <h2 className="text-2xl font-semibold mb-4 text-center">
        Votos por Hora
      </h2>
      <Bar data={data} options={options} />
    </div>
  );
}
