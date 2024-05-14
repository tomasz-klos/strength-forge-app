import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import axios from "axios";
import { useMutation } from "@tanstack/react-query";

import {
  Form,
  FormItem,
  FormLabel,
  FormField,
  FormMessage,
} from "@molecules/form";
import { Button } from "@atoms/button";
import { Input } from "@atoms/input";

interface RegisterForm {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
}

const schema = z.object({
  name: z
    .string()
    .min(3, "Name must be at least 3 characters long")
    .max(100, "Name must be at most 100 characters long"),
  email: z.string().email("Invalid email address"),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long")
    .max(100, "Password must be at most 100 characters long"),
  confirmPassword: z
    .string()
    .min(8, "Password must be at least 8 characters long")
    .max(100, "Password must be at most 100 characters long"),
});

const createUser = async (data: RegisterForm) => {
  const response = await axios.post("/api/register", data);

  console.log(response.data);
  return response.data;
};

const RegisterForm = () => {
  const form = useForm({
    resolver: zodResolver(schema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const { mutate } = useMutation({
    mutationFn: createUser,
    onSuccess: (data) => {
      console.log(data);
    },
    onError: (error) => {
      console.error(error);
    },
  });

  const onSubmit = form.handleSubmit((data) => {
    console.log(data);
    mutate(data);
  });

  return (
    <Form {...form}>
      <form className="flex flex-col gap-4" onSubmit={onSubmit}>
        <FormField
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="name">Name</FormLabel>
              <Input id="name" {...field} />
              <FormMessage>
                This will be your display name on the site
              </FormMessage>
            </FormItem>
          )}
        />
        <FormField
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="email">Email</FormLabel>
              <Input id="email" {...field} />
              <FormMessage>
                We will never share your email with anyone else
              </FormMessage>
            </FormItem>
          )}
        />
        <FormField
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="password">Password</FormLabel>
              <Input id="password" type="password" {...field} />
              <FormMessage>
                Password must be at least 8 characters long
              </FormMessage>
            </FormItem>
          )}
        />
        <FormField
          name="confirmPassword"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="confirmPassword">Confirm Password</FormLabel>
              <Input id="confirmPassword" type="password" {...field} />
              <FormMessage>Must match the password above</FormMessage>
            </FormItem>
          )}
        />
        <Button className="w-full mt-2" type="submit">
          Register
        </Button>
      </form>
    </Form>
  );
};

export default RegisterForm;
