import { useNavigate } from "react-router-dom";

import { useToast } from "@atoms/use-toast";
import { logoutUser } from "@services/auth_services";

const useLogout = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  const logout = async () => {
    try {
      await logoutUser();

      navigate("/login");
      toast({
        title: "Success",
        description: "You have successfully logged out",
      });
    } catch (error) {
      console.error(error);
      toast({
        title: "Error",
        description: "Logout failed",
        variant: "destructive",
      });
    }
  };

  return { logout };
};

export default useLogout;
