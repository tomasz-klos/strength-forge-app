import axios, { AxiosResponse } from "axios";

import { API_ROUTES } from "@config/enums";
import type { RegisterFormValues, SignInFormValues } from "@shared/form_types";

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

export const logoutUser = async () => {
  try {
    const response: AxiosResponse = await axios.post(API_ROUTES.LOGOUT);
    return response.data;
  } catch (error) {
    console.error("Error logging out:", error);
    throw error;
  }
};

export const validateToken = async () => {
  try {
    const response: AxiosResponse = await axios.get(API_ROUTES.VALIDATE_TOKEN);
    return response.data;
  } catch (error) {
    console.error("Error validating token:", error);
    throw error;
  }
};
