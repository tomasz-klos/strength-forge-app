import { Link } from "react-router-dom";

import RegisterForm from "@organisms/registerForm";
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
    <Card className="max-w-sm w-full">
      <CardHeader>
        <CardTitle className="text-lg">Create an account</CardTitle>
        <CardDescription>Sign up to get started</CardDescription>
      </CardHeader>
      <CardContent>
        <RegisterForm />
      </CardContent>
      <CardFooter className="flex items-center justify-between gap-2 mt-2">
        <CardDescription>Already have an account?</CardDescription>
        <Link className={buttonVariants({ variant: "outline" })} to="/login">
          Log in
        </Link>
      </CardFooter>
    </Card>
  );
};

export default Register;
