const BASE_URL = import.meta.env.VITE_BACKEND_URL;
const PARTICIPANTS_STORAGE = "participants-cache";

export async function getParticipants(): Promise<any[]> {
  const cache = localStorage.getItem(PARTICIPANTS_STORAGE);

  if (cache) {
    try {
      return JSON.parse(cache);
    } catch (err) {
      console.warn("Erro ao parsear cache:", err);
      localStorage.removeItem(PARTICIPANTS_STORAGE);
    }
  }

  const data = await fetch(`${BASE_URL}/api/participants`)
    .then((response) => {
      if (!response.ok) throw new Error("Failed to fetch participants");
      return response.json();
    })
    .then((parsedResponse) => {
      localStorage.setItem(
        PARTICIPANTS_STORAGE,
        JSON.stringify(parsedResponse)
      );
      return parsedResponse;
    });

  return data;
}

export async function postVote(
  selectedParticipantId: FormDataEntryValue
): Promise<void> {
  const response = await fetch(`${BASE_URL}/api/vote`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(selectedParticipantId),
  });

  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || "Erro ao registrar voto");
  }
}

export async function getResults(): Promise<Record<string, number>> {
  const response = await fetch(`${BASE_URL}/api/results`);

  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || "Erro ao buscar resultados");
  }

  const data = await response.json();

  const totalVotes = data.reduce(
    (acc: number, curr: any) => acc + curr.count,
    0
  );
  const formatted = Object.fromEntries(
    data.map((r: any) => [
      r.participant_id,
      Math.round((r.count / totalVotes) * 100),
    ])
  );

  return formatted;
}
