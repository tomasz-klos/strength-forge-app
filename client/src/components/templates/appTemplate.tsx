import { Outlet, redirect } from "react-router-dom";
import axios from "axios";

import { Toaster } from "@atoms/toaster";
import Header from "@organisms/header";

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
    const res = await axios.get("/api/auth/validate-token");
    if (res.status === 200) {
      return null;
    }
    return redirect("/login");
  } catch (error) {
    return redirect("/login");
  }
};
