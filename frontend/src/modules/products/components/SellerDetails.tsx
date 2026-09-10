import {
  Avatar,
  Box,
  Card,
  Skeleton,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import { IUser } from "../../users/interfaces/IUser";
import { userService } from "../../users/services/userService";

export const SellerDetails = ({ sellerID }: { sellerID: number }) => {
  const [seller, setseller] = useState<null | IUser>(null);
  const [error, seterror] = useState(false);
  const [loading, setloading] = useState(true);

  const fetchSeller = async () => {
    console.log("sellerID en fetchseller",sellerID)
    const response = await userService.GetUserById(sellerID);
    console.log("response en fetchseller",response)
    setloading(false);
    if (response.success && response.data) {
      setseller(response.data);
    } else {
      console.log("Error al obtener los detalles del usuario en fetchSeller");
      seterror(true);
    }
  };

  useEffect(() => {
    fetchSeller();
  }, []);

  if (loading) {
    return (
      <Card
        sx={{ width: 250, p: 2, display: "flex", alignItems: "center", gap: 2 }}
      >
        <Skeleton variant="circular" width={50} height={50} />
        <Box>
          <Skeleton variant="text" width={100} height={20} />
          <Skeleton variant="text" width={150} height={20} />
        </Box>
      </Card>
    );
  }

  if (error) {
    return (
      <Typography color="error" variant="body2">
        Error al cargar la información del vendedor.
      </Typography>
    );
  }

  return (
    <>
      {seller && (
    <Box display="flex" alignItems="center" gap={2}>
    <Avatar>{seller.Name?.charAt(0)}</Avatar>
    <Box>
      <Typography variant="subtitle1" fontWeight="bold">
        {seller.Name}
      </Typography>
      <Typography variant="body2" color="text.secondary">
        {seller.Email}
      </Typography>
    </Box>
  </Box>
      )}
    </>
  );
};
