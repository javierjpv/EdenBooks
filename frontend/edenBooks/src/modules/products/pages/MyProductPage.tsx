import { Container, Typography } from "@mui/material";
import { ProductList } from "../components/ProductList";
import { useEffect } from "react";
import { useSearchParams } from "react-router";
import { useAuthStore } from "../../users/hooks/useAuthStore";

export const MyProductPage = () => {
  const { user } = useAuthStore();
  const [searchParams, setSearchParams] = useSearchParams();

  useEffect(() => {
    const newParams = new URLSearchParams(searchParams);
    const userId = user.ID ? `${user.ID}` : "1";

    if (newParams.get("user_id") !== userId) {
      newParams.set("user_id", userId);
      setSearchParams(newParams, { replace: true }); // Evita que se agregue al historial
    }
  }, [user.ID, searchParams, setSearchParams]);

  return (
    <>
      <Container sx={{ marginTop: 22 }}>
        <Typography textAlign={"center"} variant="h4">
          My ProductPage
        </Typography>

        <ProductList />
      </Container>
    </>
  );
};
