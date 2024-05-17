import { Button, buttonVariants } from "@atoms/button";
import useLogout from "@hooks/useLogout";

const Settings = () => {
  const { logout } = useLogout();

  return (
    <section className="flex-1 flex flex-col items-center justify-center gap-4">
      <h1 className="sticky top-0 text-xl font-medium text-center">Settings</h1>

      <ul className="flex flex-col gap-2 mt-auto">
        <li>
          <Button
            className={buttonVariants({ variant: "secondary" })}
            onClick={logout}
          >
            Log Out
          </Button>
        </li>
      </ul>
    </section>
  );
};

export default Settings;
