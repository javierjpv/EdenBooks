import { IProduct } from "../interfaces/IProduct";
import {
  Avatar,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  Typography,
} from "@mui/material";
import { MoreVert } from "@mui/icons-material";
import { Link } from "react-router";
import { useAuthStore } from "../../users/hooks/useAuthStore";

export const FavoriteItem = ({ product }: { product: IProduct }) => {
  const { user } = useAuthStore();


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
      </CardActions>
    </Card>
  );
};
