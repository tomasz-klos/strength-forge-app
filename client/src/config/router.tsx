import { createBrowserRouter } from "react-router-dom";

import Home from "@pages/home";
import Register from "@pages/register";
import SignIn from "@pages/signIn";
import AppLayout from "@templates/appLayout";
import AuthLayout from "@templates/authLayout";

export const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
      {
        path: "/",
        element: <Home />,
      },
    ],
  },
  {
    element: <AuthLayout />,
    children: [
      {
        path: "/register",
        element: <Register />,
      },
      {
        path: "/login",
        element: <SignIn />,
      },
    ],
  },
]);
