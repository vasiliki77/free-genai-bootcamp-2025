
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Toaster } from "@/components/ui/toaster";
import { Toaster as Sonner } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import NavBar from "./components/layout/NavBar";

// Lazy load pages
import { lazy, Suspense } from "react";
const Dashboard = lazy(() => import("./pages/Dashboard"));
const StudyActivities = lazy(() => import("./pages/StudyActivities"));
const StudyActivity = lazy(() => import("./pages/StudyActivity"));
const Words = lazy(() => import("./pages/Words"));
const Word = lazy(() => import("./pages/Word"));
const Groups = lazy(() => import("./pages/Groups"));
const Group = lazy(() => import("./pages/Group"));
const Sessions = lazy(() => import("./pages/Sessions"));
const Session = lazy(() => import("./pages/Session"));
const Settings = lazy(() => import("./pages/Settings"));
const NotFound = lazy(() => import("./pages/NotFound"));

const queryClient = new QueryClient();

const App = () => (
  <QueryClientProvider client={queryClient}>
    <TooltipProvider>
      <BrowserRouter>
        <div className="min-h-screen bg-background">
          <NavBar />
          <main className="container pt-20">
            <Suspense
              fallback={
                <div className="flex items-center justify-center h-[calc(100vh-5rem)]">
                  <div className="animate-pulse">Loading...</div>
                </div>
              }
            >
              <Routes>
                <Route path="/" element={<Navigate to="/dashboard" replace />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/study-activities" element={<StudyActivities />} />
                <Route path="/study-activities/:id" element={<StudyActivity />} />
                <Route path="/words" element={<Words />} />
                <Route path="/words/:id" element={<Word />} />
                <Route path="/groups" element={<Groups />} />
                <Route path="/groups/:id" element={<Group />} />
                <Route path="/sessions" element={<Sessions />} />
                <Route path="/sessions/:id" element={<Session />} />
                <Route path="/settings" element={<Settings />} />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </Suspense>
          </main>
        </div>
        <Toaster />
        <Sonner />
      </BrowserRouter>
    </TooltipProvider>
  </QueryClientProvider>
);

export default App;
