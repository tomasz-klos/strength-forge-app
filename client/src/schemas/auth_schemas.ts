import { z } from "zod";

export const signInSchema = z.object({
  email: z
    .string()
    .min(3, { message: "Email must be at least 3 characters long" })
    .max(100, { message: "Email must be at most 100 characters long" })
    .email({ message: "Invalid email address" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters long" })
    .max(100, { message: "Password must be at most 100 characters long" }),
});

export const registerSchema = z
  .object({
    name: z
      .string()
      .min(3, { message: "Name must be at least 3 characters long" })
      .max(100, { message: "Name must be at most 100 characters long" }),
    email: z.string().email({ message: "Invalid email address" }),
    password: z
      .string()
      .min(8, { message: "Password must be at least 8 characters long" })
      .max(100, { message: "Password must be at most 100 characters long" }),
    confirmPassword: z
      .string()
      .min(8, { message: "Password must be at least 8 characters long" })
      .max(100, { message: "Password must be at most 100 characters long" }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  });
