import { useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
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
import { useToast } from "@atoms/use-toast";

interface SignInForm {
  email: string;
  password: string;
}

const schema = z.object({
  email: z
    .string()
    .min(3, "Name must be at least 3 characters long")
    .max(100, "Name must be at most 100 characters long"),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long")
    .max(100, "Password must be at most 100 characters long"),
});

const logInUser = async (data: SignInForm) => {
  const response = await axios.post("/api/auth/login", data);

  return response.data;
};

const SignInForm = () => {
  const navigate = useNavigate();
  const { toast } = useToast();
  const form = useForm({
    resolver: zodResolver(schema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const { mutate } = useMutation({
    mutationFn: logInUser,
    onSuccess: () => {
      navigate("/");

      toast({
        title: "Success",
        description: "You have successfully signed in",
      });
    },
    onError: (error) => {
      console.error(error);

      toast({
        title: "Error",
        description: "Invalid email or password",
        variant: "destructive",
      });
    },
  });

  const onSubmit = form.handleSubmit((data) => {
    mutate(data);
  });

  return (
    <Form {...form}>
      <form className="flex flex-col gap-4" onSubmit={onSubmit}>
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
        <Button className="w-full mt-2" type="submit">
          Sign in
        </Button>
      </form>
    </Form>
  );
};

export default SignInForm;
