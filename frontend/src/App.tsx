import { useQuery } from '@tanstack/react-query';
import { fetchHealth } from './api/health';
import './App.css';

function App() {
  const { data, isLoading, isError } = useQuery({
    queryKey: ['health'],
    queryFn: fetchHealth,
  });

  if (isLoading) return <div className="p-4 text-gray-600">Loading...</div>;
  if (isError) return <div className="p-4 text-red-500">Error</div>;

  return (
    <div className="min-h-screen flex items-center justify-center bg-slate-100">
      <div className="bg-white shadow rounded-xl p-6">
        <h1 className="text-2xl font-bold mb-2">Go + React Monorepo</h1>
        <p className="text-gray-700">
          Backend health: <span className="font-mono">{data?.status}</span>
        </p>
      </div>
    </div>
  );
}

export default App
