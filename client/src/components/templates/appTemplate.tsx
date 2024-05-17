import { Outlet } from "react-router-dom";

import { Toaster } from "@atoms/toaster";

const AppTemplate = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1">
        <Outlet />
      </main>
      <Toaster />
    </div>
  );
};

export default AppTemplate;
