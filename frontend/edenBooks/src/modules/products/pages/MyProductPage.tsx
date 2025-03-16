import { Container } from "@mui/material";
import { ProductList } from "../components/ProductList";
import { ProductFilters } from "../components/ProductFilters";

export const MyProductPage = () => {


  return (
    <>
      <Container sx={{ marginTop: 22 }}>
        <h1>My ProductPage</h1>
        <ProductFilters/>
        <ProductList  />
      </Container>
    </>
  );
};
