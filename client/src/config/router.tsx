import { createBrowserRouter } from "react-router-dom";
import Home from "@pages/home";
import Register from "@pages/register";
import SignIn from "@pages/signIn";
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
