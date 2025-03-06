import { Navigate, Outlet } from "react-router";
import { useAuthStore } from "../modules/users/hooks/useAuthStore";

export const ProtectedRoute = () => {
     const { user} = useAuthStore();
    return user?.userState === "AUTHENTICATED" ? <Outlet /> : <Navigate to="/auth" />;
  };