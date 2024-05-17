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
import { signInschema } from "@schemas/auth_schemas";
import { signInUser } from "@services/auth_services";

import { SignInFormValues } from "@shared/form_types";

const SignInForm: React.FC = () => {
  const { form, onSubmit } = useAuthForm<SignInFormValues>({
    schema: signInschema,
    mutationFn: signInUser,
    onSuccessRedirect: "/",
    successMessage: "You have successfully signed in",
    errorMessage: "Invalid email or password",
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
