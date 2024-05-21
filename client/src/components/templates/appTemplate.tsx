import { Outlet, redirect, useLoaderData } from "react-router-dom";

import { Toaster } from "@atoms/toaster";
import Header from "@organisms/header";
import { validateToken } from "@services/auth_services";
import { User } from "@shared/user_types";

const AppTemplate: React.FC = () => {
  const user = useLoaderData() as User;

  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1 flex flex-col p-4 overflow-y-auto">
        <Outlet context={{ user }} />
      </main>
      <Header />
      <Toaster />
    </div>
  );
};

export default AppTemplate;

export const appTemplateLoader = async () => {
  try {
    const res = await validateToken();

    return res.data;
  } catch (error) {
    return redirect("/login");
  }
};
