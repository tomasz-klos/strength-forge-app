import { Outlet, redirect } from "react-router-dom";

import { Toaster } from "@atoms/toaster";
import Header from "@organisms/header";
import { validateToken } from "@services/auth_services";

const AppTemplate = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1 flex flex-col p-4 overflow-y-auto">
        <Outlet />
      </main>
      <Header />
      <Toaster />
    </div>
  );
};

export default AppTemplate;

export const appTemplateLoader = async () => {
  try {
    await validateToken();

    return null;
  } catch (error) {
    return redirect("/login");
  }
};
