import { Link } from "react-router-dom";

import SignInForm from "@organisms/signInForm";
import {
  Card,
  CardContent,
  CardHeader,
  CardFooter,
  CardTitle,
  CardDescription,
} from "@atoms/card";
import { buttonVariants } from "@atoms/button";

const Register = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-lg">Welcome to App</CardTitle>
        <CardDescription>Sign in to get started</CardDescription>
      </CardHeader>
      <CardContent>
        <SignInForm />
      </CardContent>
      <CardFooter className="flex items-center justify-between gap-2 mt-2">
        <CardDescription>Don't have an account?</CardDescription>
        <Link className={buttonVariants({ variant: "outline" })} to="/register">
          Register
        </Link>
      </CardFooter>
    </Card>
  );
};

export default Register;
