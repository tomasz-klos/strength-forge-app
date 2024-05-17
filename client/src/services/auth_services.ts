import axios from "axios";

import { RegisterFormValues, SignInFormValues } from "@shared/form_types";
import { API_ROUTES } from "@config/enums";

export const signInUser = async (data: SignInFormValues) => {
  const response = await axios.post(API_ROUTES.LOGIN, data);
  return response.data;
};

export const registerUser = async (data: RegisterFormValues) => {
  const response = await axios.post(API_ROUTES.REGISTER, data);
  return response.data;
};
