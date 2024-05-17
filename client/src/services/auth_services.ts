import axios from "axios";

import { RegisterFormValues, SignInFormValues } from "@shared/form_types";

export const signInUser = async (data: SignInFormValues) => {
  const response = await axios.post("/api/auth/login", data);
  return response.data;
};

export const registerUser = async (data: RegisterFormValues) => {
  const response = await axios.post("/api/auth/register", data);
  return response.data;
};
