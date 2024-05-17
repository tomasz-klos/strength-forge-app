import {
  Form,
  FormItem,
  FormLabel,
  FormField,
  FormMessage,
} from "@molecules/form";
import { Button } from "@atoms/button";
import { Input } from "@atoms/input";

import useAuthForm from "@hooks/useAuthForm";
import { registerSchema } from "@schemas/auth_schemas";
import { registerUser } from "@services/auth_services";

import { RegisterFormValues } from "@shared/form_types";

const RegisterForm: React.FC = () => {
  const { form, onSubmit } = useAuthForm<RegisterFormValues>({
    schema: registerSchema,
    mutationFn: registerUser,
    onSuccessRedirect: "/",
    successMessage: "You have successfully registered",
    errorMessage: "Registration failed",
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
