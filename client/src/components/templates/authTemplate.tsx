import { Outlet, redirect } from "react-router-dom";

import { Toaster } from "@atoms/toaster";
import { validateToken } from "@services/auth_services";

const AuthTemplate: React.FC = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1 flex items-center justify-center px-4">
        <Outlet />
      </main>
      <Toaster />
    </div>
  );
};

export default AuthTemplate;

export const authTemplateLoader = async (): Promise<null | Response> => {
  try {
    await validateToken();

    return redirect("/");
  } catch (error) {
    return null;
  }
};
