import { createBrowserRouter } from "react-router-dom";

import Home from "@pages/home";
import Profile from "@pages/profile";
import Register from "@pages/register";
import SignIn from "@pages/signIn";
import Settings from "@pages/settings";
import AppTemplate, { appTemplateLoader } from "@templates/appTemplate";
import AuthTemplate, { authTemplateLoader } from "@templates/authTemplate";

export const router = createBrowserRouter([
  {
    element: <AppTemplate />,
    loader: appTemplateLoader,
    children: [
      {
        path: "/",
        element: <Home />,
      },
      {
        path: "/profile",
        element: <Profile />,
      },
      {
        path: "/settings",
        element: <Settings />,
      },
    ],
  },
  {
    element: <AuthTemplate />,
    loader: authTemplateLoader,
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
