import http from "k6/http";
import { check } from "k6";

const BASE_URL = __ENV.BASE_URL || "http://localhost:8080";

export let options = {
  scenarios: {
    valid_votes: {
      executor: "constant-arrival-rate",
      rate: 1250,
      timeUnit: "1s",
      duration: "30s",
      preAllocatedVUs: 100,
      maxVUs: 500,
    },
  },
};

export function setup() {
  const res = http.get(`${BASE_URL}/api/participants?isNominated=true`);

  if (res.status !== 200) {
    throw new Error(`Erro ao buscar participantes: status ${res.status}`);
  }

  const participants = JSON.parse(res.body);
  if (!Array.isArray(participants) || participants.length === 0) {
    throw new Error("Lista de participantes vazia ou inválida.");
  }

  return participants.map((p) => p.id);
}

export default function (participants) {
  const selected =
    participants[Math.floor(Math.random() * participants.length)];

  const payload = JSON.stringify({
    participant_id: selected,
  });

  const headers = { "Content-Type": "application/json" };

  const res = http.post(`${BASE_URL}/api/vote`, payload, { headers });

  check(res, {
    "status is 200": (r) => r.status === 200,
  });
}