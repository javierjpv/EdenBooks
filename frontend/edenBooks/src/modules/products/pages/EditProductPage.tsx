import { Container } from "@mui/material";
import { ProductForm } from "../components/ProductForm";
import { useParams } from "react-router";

export const EditProductPage = () => {
  const { id } = useParams();
  const idNumber = Number(id);

  return (
    <>
      <Container maxWidth="sm" sx={{ marginTop: 29 }}>
        <ProductForm id={idNumber} />
      </Container>
    </>
  );
};
