import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useToast } from "@atoms/use-toast";

const useLogout = () => {
  const navigate = useNavigate();
  const { toast } = useToast();

  const logout = async () => {
    try {
      await axios.post("/api/auth/logout");
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
