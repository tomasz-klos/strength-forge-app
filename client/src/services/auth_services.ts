import axios, { AxiosResponse } from "axios";

import { RegisterFormValues, SignInFormValues } from "@shared/form_types";
import { API_ROUTES } from "@config/enums";

export const signInUser = async (data: SignInFormValues) => {
  try {
    const response: AxiosResponse = await axios.post(API_ROUTES.LOGIN, data);
    return response.data;
  } catch (error) {
    console.error("Error signing in:", error);
    throw error;
  }
};

export const registerUser = async (data: RegisterFormValues) => {
  try {
    const response: AxiosResponse = await axios.post(API_ROUTES.REGISTER, data);
    return response.data;
  } catch (error) {
    console.error("Error registering user:", error);
    throw error;
  }
};
