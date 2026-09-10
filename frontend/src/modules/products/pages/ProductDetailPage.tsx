import { useNavigate, useParams } from "react-router";
import { useEffect, useState } from "react";
import { productService } from "../services/productService";
import { IProduct } from "../interfaces/IProduct";
import { Box, Button, Container, Skeleton } from "@mui/material";
import { ProductDetail } from "../components/ProductDetail";
import { ReviewList } from "../../reviews/components/ReviewList";
import { useAuthStore } from "../../users/hooks/useAuthStore";
import { useDispatch } from "react-redux";
import { updateProduct } from "../../../Store/checkout/checkoutSlice";
export const ProductDetailPage = () => {
  const { user } = useAuthStore();
  const { id } = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const [error, seterror] = useState(false);
  const [productError, setproductError] = useState<boolean>(false);
  const [loading, setloading] = useState<boolean>(true);
  const [product, setproduct] = useState<IProduct | null>(null);

  const getProduct = async (): Promise<void> => {
    const response = await productService.GetProductById(Number(id));
    setTimeout(() => {
      setloading(false);
    }, 200);
    if (response.success && response.data) {
      setproductError(true);
      const fetchedProduct: IProduct = response.data;
      console.log("Product fetched", fetchedProduct);
      setproduct(fetchedProduct);
      return;
    }
    setproductError(true);
    console.log("Error al obntener el producto en productDetailPage");
  };

  const checkProduct = (product: IProduct): boolean => {
    console.log("CHECKING PRODUCT: ", product);
    if (!product.ImageURL || !product.ID) {
      return false;
    }
    if (
      product.Description.trim() === "" ||
      product.ImageURL.trim() === "" ||
      product.Price <= 0 ||
      product.Name.trim() === "" ||
      product.UserID <= 0 ||
      product.CategoryID <= 0 ||
      product.ID <= 0 ||
      product.Sold === true
    ) {
      return false;
    }
    return true;
  };

  const handlePurchase = () => {
    if (!product) {
      console.log("El producto aun no se ha cargado desde la funcion getByID");
      return;
    }
    const prod: IProduct = { ...product };
    if (checkProduct(prod)) {
      dispatch(updateProduct(prod));
      navigate("/checkout/shipping");
      console.log("TODO HA IDO BIEN, PRODUCT:", product);
      return;
    }
    //Contruir pagina de error generica
    console.log("ERROR AL DIRIGIRSE A SHIPPING");
    console.log("VALOR DE PRODUCT EN EL ERROR", product);
    seterror(true);
    setTimeout(() => {
      seterror(false);
    }, 3000);
  };

  useEffect(() => {
    getProduct();
  }, [id]);

  return (
    <Container sx={{ marginTop: 22 }} maxWidth="md">
      {loading ? (
        <>
          <Skeleton variant="text" width="60%" height={40} />
          <Skeleton
            variant="rectangular"
            width="100%"
            height={300}
            sx={{ mt: 2 }}
          />
          <Skeleton variant="text" width="40%" sx={{ mt: 2 }} />
          <Skeleton variant="text" width="80%" sx={{ mt: 1 }} />
        </>
      ) : !product && productError ? (
        <p>Ha ocurrido un error en ProductDetail</p>
      ) : (
        <>
          {product && (
            <>
              <ProductDetail product={product} />
              <ReviewList />

              {product &&
                user.userState === "AUTHENTICATED" &&
                Number(user.ID) !== product.UserID &&
                !product.Sold && (
                  <Box
                    marginTop={3}
                    display={"flex"}
                    flexDirection={"row"}
                    justifyContent={"center"}
                  >
                    <Button
                      onClick={handlePurchase}
                      fullWidth
                      variant="contained"
                      sx={{ marginBottom: 3, borderRadius: 4, padding: 1.3 }}
                    >
                      Comprar
                    </Button>
                  </Box>
                )}
            </>
          )}
        </>
      )}
    </Container>
  );
};
