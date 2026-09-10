import { IProduct } from "../interfaces/IProduct";
import {
  Alert,
  Avatar,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  CircularProgress,
  IconButton,
  Snackbar,
  Typography,
} from "@mui/material";
import { Delete, Edit, Favorite, MoreVert } from "@mui/icons-material";
import { productService } from "../services/productService";
import { useState } from "react";
import { Link, useNavigate } from "react-router";
import { useAuthStore } from "../../users/hooks/useAuthStore";
interface ProductItemProps{
  product:IProduct
  fetchProducts:()=>void
}
export const ProductItem = ({ product,fetchProducts }:ProductItemProps) => {
  const { user } = useAuthStore();
  const [loadingDelete, setloadingDelete] = useState<boolean>(false);
  const [isFavorite, setIsFavorite] = useState(product.isFavorite);
  const [openSnackbar, setOpenSnackbar] = useState<boolean>(false);
  const navigate = useNavigate();

  const handleEdit = (): void => {
    if (product?.ID) {
      navigate(`/products/edit/${product.ID}`);
    } else {
      console.log("No se ha accedido correctamente a la edicion");
    }
  };
  const handleDelete = async (id: number): Promise<void> => {
    try {
      console.log("Iniciando eliminacion");
      setloadingDelete(true);
      await productService.DeleteProduct(id);
      console.log("eliminacion completada");
      setTimeout(() => {
        setloadingDelete(false);
      }, 700);
      fetchProducts()
    } catch (error) {
      setloadingDelete(false);
      console.log("Error al eliminar un producto", error);
    }
  };
  const handleFavorite = async () => {
    if (!product || !product.ID) {
      console.log("El producto no existe");
      return;
    }
    try {
      if (isFavorite) {
        await productService.DeleteFromFavorite(product.ID);
        setIsFavorite(false);
      } else {
        await productService.AddToFavorite(product.ID);
        setIsFavorite(true);
        setOpenSnackbar(true);
      }
    } catch (error) {
      console.log("Error al actualizar favorito", error);
    }
  };

  return (
  <>
   <Card>
      <Link
        to={`/products/${product.ID}`}
        style={{ textDecoration: "none", color: "inherit" }}
      >
        <CardHeader
          avatar={
            <Avatar sx={{ bgcolor: "red" }} aria-label="product">
              {product.Name[0]}
            </Avatar>
          }
          action={
            <IconButton aria-label="settings">
              <MoreVert />
            </IconButton>
          }
          title={product.Name}
        />

        <CardMedia
          component="img"
          height="330"
          image={product.ImageURL ?? "/418eyXxdCsL._SY445_SX342_.jpg"}
          // image={product.ImageURL|| "/418eyXxdCsL._SY445_SX342_.jpg"}
          alt={product.Name}
        />

        <CardContent>
          <Typography variant="body2" sx={{ color: "text.secondary" }}>
            {product.Name}
          </Typography>

          <Typography variant="body2" sx={{ color: "text.secondary" }}>
            Precio: {product.Price} €
          </Typography>

          {user?.userState === "AUTHENTICATED" &&
            Number(user?.ID) !== product.UserID && (
              <>
                {product.Sold ? (
                  <Typography variant="body1" gutterBottom color="error">
                    VENDIDO
                  </Typography>
                ) : (
                  <Typography variant="body1" gutterBottom color="warning">
                    DISPONIBLE
                  </Typography>
                )}
              </>
            )}
        </CardContent>
      </Link>

      <CardActions disableSpacing>
        {user?.userState === "AUTHENTICATED" &&
          Number(user?.ID) === product.UserID && (
            <>
              <IconButton onClick={handleEdit} aria-label="share">
                <Edit />
              </IconButton>
              <IconButton
                disabled={loadingDelete}
                onClick={() => product.ID && handleDelete(product.ID)}
                aria-label="delete"
                style={{ display: "ID" in product ? "inline-flex" : "none" }}
              >
                {loadingDelete ? <CircularProgress size={22} /> : <Delete />}
              </IconButton>
            </>
          )}
        {user?.userState === "AUTHENTICATED" &&
          Number(user?.ID) !== product.UserID && (
            <>
              {!product.Sold && (
                <>
                        {!product.Sold && (
                          <>
                              <IconButton
                                color={isFavorite?"error":"default"}
                                onClick={handleFavorite}
                                aria-label="add to favorites"
                              >
                                <Favorite />
                              </IconButton>
                          </>
                        )}
                </>
              )}
            </>
          )}
      </CardActions>
    </Card>
          {/* Notificación Snackbar */}
          <Snackbar
        open={openSnackbar}
        autoHideDuration={3000}
        onClose={() => setOpenSnackbar(false)}
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <Alert onClose={() => setOpenSnackbar(false)} severity="info">
          ¡Producto añadido a favoritos!
        </Alert>
      </Snackbar>
  </>
  );
};
