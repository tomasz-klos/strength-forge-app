import { Button, buttonVariants } from "@atoms/button";

const Home: React.FC = () => {
  return (
    <>
      <h1 className="text-lg font-medium text-blue-500">StrengthForge App</h1>
      <Button className={buttonVariants({ variant: "outline" })}>
        Click me
      </Button>
    </>
  );
};

export default Home;
