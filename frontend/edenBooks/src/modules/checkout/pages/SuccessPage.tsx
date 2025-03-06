
import { Button, Container, Typography, Box } from "@mui/material";
import { useDispatch } from "react-redux";
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import { useEffect } from "react";
import { useNavigate } from "react-router";
import { resetCheckout } from "../../../Store/checkout/checkoutSlice";

export const SuccessPage = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  

  useEffect(() => {
    dispatch(resetCheckout());
  }, []);

  return (
    <Container maxWidth="sm" sx={{ textAlign: "center", mt: 8 }}>
      <CheckCircleIcon sx={{ fontSize: 64, color: "success.main", mb: 2 }} />
      <Typography variant="h4" gutterBottom>
        ¡Pago completado con éxito!
      </Typography>
      <Typography variant="body1" >
        Gracias por tu compra. Hemos enviado un correo electrónico con los detalles de tu pedido.
      </Typography>
      <Box sx={{ mt: 4 }}>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={() => navigate("/")}
          sx={{ mr: 2 }}
        >
          Volver a la tienda
        </Button>
        <Button 
          variant="outlined" 
          onClick={() => navigate("/orders")}
        >
          Ver mis pedidos
        </Button>
      </Box>
    </Container>
  );
};