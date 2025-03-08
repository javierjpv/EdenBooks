import { IProduct } from "../interfaces/IProduct";
import {
  Avatar,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  CircularProgress,
  IconButton,
  Typography,
} from "@mui/material";
import { Delete, Edit, Favorite, MoreVert } from "@mui/icons-material";
import { productService } from "../services/productService";
import { useState } from "react";
import { Link, useNavigate } from "react-router";
import { useAuthStore } from "../../users/hooks/useAuthStore";

export const ProductItem = ({ product }: { product: IProduct }) => {
  const { user } = useAuthStore();
  const [loadingDelete, setloadingDelete] = useState(false);
  const [isFavorite, setIsFavorite] = useState(product.is_favorite);
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
      }
    } catch (error) {
      console.log("Error al actualizar favorito", error);
    }
  };

  return (
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
  );
};
