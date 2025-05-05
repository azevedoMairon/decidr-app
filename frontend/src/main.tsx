import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import Index from "./pages/Index.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import VotingPage from "./pages/VotingPage.tsx";
import { AnimatePresence } from "framer-motion";
import { Toaster } from "react-hot-toast";
import Bracket from "./assets/Bracket.tsx";
import ParticipantProvider from "./contexts/ParticipantContext.tsx";
import VoteProvider from "./contexts/VoteContext.tsx";
import VoteStats from "./pages/VoteStats.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ParticipantProvider>
      <VoteProvider>
        <Toaster position="top-right" toastOptions={{ duration: 3000 }} />
        <BrowserRouter>
          <AnimatePresence>
            <Routes>
              <Route path="/" element={<Index />} />
              <Route path="/voting-page" element={<VotingPage />} />
              <Route path="/stats" element={<VoteStats />} />
            </Routes>
            <Bracket />
          </AnimatePresence>
        </BrowserRouter>
      </VoteProvider>
    </ParticipantProvider>
  </StrictMode>
);
