import { useNavigate } from "react-router-dom";

import { useToast } from "@atoms/use-toast";
import { logoutUser } from "@services/auth_services";

interface UseLogout {
  logout: () => Promise<void>;
}

const useLogout = (): UseLogout => {
  const navigate = useNavigate();
  const { toast } = useToast();

  const handleSuccess = () => {
    navigate("/login");
    toast({
      title: "Success",
      description: "You have successfully logged out",
    });
  };

  const handleError = (error: unknown) => {
    console.error("Logout error:", error);
    toast({
      title: "Error",
      description: "Logout failed",
      variant: "destructive",
    });
  };

  const logout = async (): Promise<void> => {
    try {
      await logoutUser();

      handleSuccess();
    } catch (error) {
      handleError(error);
    }
  };

  return { logout };
};

export default useLogout;
