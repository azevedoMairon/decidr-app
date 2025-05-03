import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import VotingPage from "./pages/VotingPage.tsx";
import { AnimatePresence } from "framer-motion";
import { Toaster } from "react-hot-toast";
import Bracket from "./components/Bracket.tsx";
import ParticipantProvider from "./contexts/ParticipantContext.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ParticipantProvider>
      <Toaster position="top-right" toastOptions={{ duration: 3000 }}></Toaster>
      <BrowserRouter>
        <AnimatePresence>
          <Routes>
            <Route path="/" element={<App />} />
            <Route path="/voting-page" element={<VotingPage />} />
          </Routes>
          <Bracket />
        </AnimatePresence>
      </BrowserRouter>
    </ParticipantProvider>
  </StrictMode>
);
