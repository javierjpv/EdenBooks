import { Container } from "@mui/material";
import { ProductList } from "../components/ProductList";

export const MyProductPage = () => {


  return (
    <>
      <Container sx={{ marginTop: 22 }}>
        <h1>My ProductPage</h1>
        <ProductList  />
      </Container>
    </>
  );
};
