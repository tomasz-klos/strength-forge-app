import { Outlet } from "react-router-dom";

const AppLayout: React.FC = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1 flex items-center justify-center px-4">
        <Outlet />
      </main>
    </div>
  );
};

export default AppLayout;
