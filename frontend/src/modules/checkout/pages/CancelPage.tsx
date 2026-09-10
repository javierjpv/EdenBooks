import { Button, Container, Typography, Box } from "@mui/material";
import ErrorIcon from '@mui/icons-material/Error';
import { useNavigate } from "react-router";
import { useDispatch } from "react-redux";
import { resetCheckout } from "../../../Store/checkout/checkoutSlice";
import { useEffect } from "react";

export const CancelPage = () => {
  const navigate = useNavigate();


  const dispatch = useDispatch();
    useEffect(() => {
      dispatch(resetCheckout());
    }, []);

  return (
    <Container maxWidth="sm" sx={{ textAlign: "center", mt: 8 }}>
      <ErrorIcon sx={{ fontSize: 64, color: "error.main", mb: 2 }} />
      <Typography variant="h4" gutterBottom>
        Pago cancelado
      </Typography>
      <Typography variant="body1" >
        El proceso de pago ha sido cancelado. No se ha realizado ningún cargo a tu cuenta.
      </Typography>
      <Box sx={{ mt: 4 }}>
        <Button 
          variant="contained" 
          color="primary" 
          onClick={() => navigate("/checkout/payment")}
          sx={{ mr: 2 }}
        >
          Intentar de nuevo
        </Button>
        <Button 
          variant="outlined" 
          onClick={() => navigate("/")}
        >
          Volver a la tienda
        </Button>
      </Box>
    </Container>
  );
};