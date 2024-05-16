import { createBrowserRouter } from "react-router-dom";

import Home from "@pages/home";
import Register from "@pages/register";
import SignIn from "@pages/signIn";
import AppTemplate from "@templates/appTemplate";
import AuthTemplate from "@templates/authTemplate";

export const router = createBrowserRouter([
  {
    element: <AppTemplate />,
    children: [
      {
        path: "/",
        element: <Home />,
      },
    ],
  },
  {
    element: <AuthTemplate />,
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
