import { IProduct } from "../interfaces/IProduct";
import {
  // Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  CircularProgress,
  Container,
  IconButton,
  Skeleton,
  Typography,
} from "@mui/material";
import {
  AccessTime,
  ArrowBackIosNew,
  Delete,
  Edit,
  Favorite,
  LocalShipping,
  MoreVert,
} from "@mui/icons-material";
import { productService } from "../services/productService";
import { useState } from "react";
import { useNavigate } from "react-router";
import { useAuthStore } from "../../users/hooks/useAuthStore";
import { SellerDetails } from "./SellerDetails";
import { chatService } from "../../chats/services/chatService";

export const ProductDetail = ({ product }: { product: IProduct | null }) => {
  const { user } = useAuthStore();
  const [loadingDelete, setloadingDelete] = useState(false);
  const navigate = useNavigate();
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
  const handleEdit = (): void => {
    if (product?.ID) {
      navigate(`/products/edit/${product.ID}`);
    } else {
      console.log("No se ha accedido correctamente a la edicion");
    }
  };
  const handleGoBack = () => {
    navigate(-1); // Esto te lleva a la página anterior
  };

  const createChat = async (): Promise<number | null> => {
    if (!user.ID || !product) {
      return null;
    }
    console.log("USUARIOS", Number(user.ID), product.UserID);
    const response = await chatService.CreateChat({
      userIDs: [Number(user.ID), product.UserID],
    });
    if (!response.success) {
      console.log("Error al crear");
      return null;
    }
    if (!response.data) {
      console.log("no contiene datos");
      return null;
    }
    console.log("RESPONSE", response);
    return response.data.ID;
  };
  const handleChat = async (): Promise<void> => {
    if (!product) {
      console.log("Error en handleChat");
    }
    const chatID = await createChat();
    if (chatID != null) {
      console.log("navegacion hacia chat con id: ", chatID);
      navigate(`/chats/${chatID}`);
    }
  };

  const handleFavorite = async () => {
    if (product && product.ID) {
      console.log("añadido a favoritos");
      try {
        await productService.AddToFavorite(product.ID);
      } catch (error) {
        console.log("error al añadir a favorito");
      } finally {
      }
    } else {
      console.log("no existe aun el product");
    }
  };

  return (
    <>
      {product == null ? (
        <Card sx={{ display: "flex", flexDirection: "column", mb: 2 }}>
          <Skeleton variant="rectangular" width="100%" height={300} />
          <CardContent>
            <Skeleton variant="text" width="60%" />
            <Skeleton variant="text" width="40%" />
          </CardContent>
          <CardActions disableSpacing>
            <Skeleton variant="circular" width={40} height={40} />
            <Skeleton
              variant="circular"
              width={40}
              height={40}
              sx={{ ml: 1 }}
            />
            <Skeleton
              variant="circular"
              width={40}
              height={40}
              sx={{ ml: 1 }}
            />
            <Skeleton
              variant="circular"
              width={40}
              height={40}
              sx={{ ml: 1 }}
            />
          </CardActions>
        </Card>
      ) : (
        <Container maxWidth="sm">
          <Button
            onClick={handleGoBack}
            startIcon={<ArrowBackIosNew />}
            sx={{ marginBottom: 3 }}
          >
            Volver
          </Button>
          <Card>
            <CardHeader
              action={
                <IconButton aria-label="settings">
                  <MoreVert />
                </IconButton>
              }
              subheader={
                product &&
                user?.userState === "AUTHENTICATED" &&
                Number(user.ID) !== product.UserID &&
                !product.Sold && (
                  <>
                    {/* Crear componente */}
                    {/* Al componennte que crees pasarle el product para asi obtener el userID */}
                    <Box
                      marginTop={4}
                      display={"flex"}
                      justifyContent={"space-between"}
                      alignContent={"center"}
                      gap={2}
                    >
                      {/* <p>Vendedor ID: {product?.UserID}</p> */}
                      <SellerDetails sellerID={product.UserID}></SellerDetails>

                      <Box>
                        {!product.Sold && (
                          <IconButton
                            onClick={handleFavorite}
                            aria-label="add to favorites"
                          >
                            <Favorite />
                          </IconButton>
                        )}
                        <Button
                          onClick={handleChat}
                          variant="outlined"
                          sx={{ borderRadius: 6, padding: 0.5 }}
                        >
                          Chat
                        </Button>
                      </Box>
                    </Box>
                    {/* Crear componente */}
                  </>
                )
              }
            />

            <CardMedia
              component="img"
              height="330"
              image={product.ImageURL ?? "/418eyXxdCsL._SY445_SX342_.jpg"}
              alt={product.Name}
            />

            <CardContent>
              <Typography variant="body1" sx={{ color: "text.secondary" }}>
                Precio: {product.Price} €
              </Typography>
              <Typography variant="body1" sx={{ color: "text.secondary" }}>
                {product.Name}
              </Typography>

              <Typography variant="body2" sx={{ color: "text.secondary" }}>
                <></> Envio disponible
              </Typography>

              {/* Aqui colocar el avatar del vendedor junto con el chat a la derecha  */}
              <Typography variant="body2" sx={{ color: "text.secondary" }}>
                Vendedor ID: {product.UserID}
              </Typography>

              <Typography variant="body2" sx={{ color: "text.secondary" }}>
                Detalles del producto: {product.Description}
              </Typography>

              {/* Información de Entrega */}
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: 1,
                  mt: 2,
                  p: 2,
                  border: "1px solid #ddd",
                  borderRadius: "8px",
                  bgcolor: "#f9f9f9",
                }}
              >
                <LocalShipping color="primary" />
                <Typography variant="body2" sx={{ color: "text.primary" }}>
                  Entrega en 3-7 días hábiles
                </Typography>
                <AccessTime color="action" />
              </Box>
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
                      style={{
                        display: "ID" in product ? "inline-flex" : "none",
                      }}
                    >
                      {loadingDelete ? (
                        <CircularProgress size={22} />
                      ) : (
                        <Delete />
                      )}
                    </IconButton>
                  </>
                )}
            </CardActions>
          </Card>
        </Container>
      )}
    </>
  );
};
