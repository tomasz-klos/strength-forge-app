import { Outlet, redirect } from "react-router-dom";
import axios from "axios";

import { Toaster } from "@atoms/toaster";
import { API_ROUTES } from "@config/enums";

const AppTemplate = () => {
  return (
    <div className="flex flex-col min-h-screen bg-zinc-950 text-zinc-50">
      <main className="flex-1 flex items-center justify-center px-4">
        <Outlet />
      </main>
      <Toaster />
    </div>
  );
};

export default AppTemplate;

export const authTemplateLoader = async () => {
  try {
    const res = await axios.get(API_ROUTES.VALIDATE_TOKEN);
    if (res.status === 200) {
      return redirect("/");
    }
    return null;
  } catch (error) {
    return null;
  }
};
