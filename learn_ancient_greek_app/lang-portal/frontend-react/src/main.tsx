import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import TranslationComponent from './components/ui/TranslationComponent'

createRoot(document.getElementById("root")!).render(<App />);
