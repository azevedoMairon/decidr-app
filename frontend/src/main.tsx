import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import VotingPage from "./pages/VotingPage.tsx";
import { AnimatePresence } from "framer-motion";
import Bracket from "./components/Bracket.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <AnimatePresence mode="wait">
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/voting-page" element={<VotingPage />} />
        </Routes>
        <Bracket />
      </AnimatePresence>
    </BrowserRouter>
  </StrictMode>
);
