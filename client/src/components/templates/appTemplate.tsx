import { Outlet } from "react-router-dom";

const AppTemplate = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1">
        <Outlet />
      </main>
    </div>
  );
};

export default AppTemplate;
