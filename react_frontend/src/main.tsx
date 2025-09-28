import React, { createContext } from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import App from './App';
import './index.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const WindowWidthContext = createContext<number | null>(null)
const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root')!).render(
  // Mounts components twice to check sidefx, only in dev
  <React.StrictMode>
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <WindowWidthContext.Provider value={512}>
          <App />
        </WindowWidthContext.Provider>
      </QueryClientProvider>
    </BrowserRouter>
  </React.StrictMode>
);