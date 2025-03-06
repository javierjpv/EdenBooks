import { Container } from "@mui/material";
import { ProductList } from "../components/ProductList";


export const ProductPage = () => {


  return (
    <>
      <Container sx={{ marginTop: 22 }}>
        <h1>My ProductPage</h1>
        <ProductList  />
      </Container>
    </>
  );
};
