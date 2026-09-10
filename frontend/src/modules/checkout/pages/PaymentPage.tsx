import {
  Typography,
  Button,
  Box,
  Modal,
  Alert,
  CircularProgress,
  Card,
  CardMedia,
} from "@mui/material";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router";
import { loadStripe } from "@stripe/stripe-js";
import axios, { AxiosError } from "axios";
import { STRIPE_CONFIG } from "../../../config/stripe";
import { useAuthStore } from "../../users/hooks/useAuthStore";
import { ArrowBackIosNew } from "@mui/icons-material";

const BASE_URL = "http://localhost:6969/stripe";
// Cargamos Stripe fuera del componente para evitar múltiples instancias
const stripePromise = loadStripe(STRIPE_CONFIG.publishableKey);

export const PaymentPage = () => {
  const { user } = useAuthStore();
  const navigate = useNavigate();
  const { shipping, product,carrier} = useSelector((state: any) => state.checkout);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const userID = Number(user.ID);
  const carrierID = Number(carrier.ID);
  const productID = Number(product.ID);

  const requestBody = {
    product,
    shipping,
    userID,
    productID,
    carrierID,
    successUrl: `${window.location.origin}/checkout/success`,
    cancelUrl: `${window.location.origin}/checkout/cancel`,
  };

  const handlePayment = async () => {
    console.log("SHIPPING", shipping);
    console.log("PRODUCT", product);
    console.log("VALOR DE REQUEST BODY", requestBody);
    try {
      setLoading(true);
      setError(null);

      const stripe = await stripePromise;
      if (!stripe) throw new Error("No se pudo cargar Stripe");

      const response = await axios.post(BASE_URL, requestBody);

      const { sessionId } = response.data;
      const result = await stripe.redirectToCheckout({ sessionId });

      if (result.error) {
        throw new Error(result.error.message ?? "Error en el proceso de pago");
      }
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      setError(axiosError.request?.response);
      console.error("Error en el pago:", axiosError.request?.response);
    } finally {
      setLoading(false);
    }
  };

  const modalStyle = {
    position: "absolute",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: 400,
    bgcolor: "background.paper",
    border: "2px solid #000",
    boxShadow: 24,
    p: 4,
  };

  // Verificar si tenemos la información de envío
  if (!shipping || !shipping.street) {
    navigate("/checkout/shipping");
    return null;
  }
  const handleGoBack = () => {
    navigate(-1);
  };

  return (
    <div>
      <Modal
        open={true}
        aria-labelledby="payment-modal-title"
        aria-describedby="payment-modal-description"
      >
        <Box sx={modalStyle}>
          <Button
            onClick={handleGoBack}
            startIcon={<ArrowBackIosNew />}
            sx={{ marginBottom: 3 }}
          >
            Volver
          </Button>
          <Typography id="payment-modal-title" variant="h6" component="h2">
            Checkout Information
          </Typography>

          <Box sx={{ mt: 2, mb: 3 }}>
            <Typography variant="body1" gutterBottom>
              Shipping to: {shipping.street} {shipping.number}, {shipping.city}
            </Typography>
            <Typography variant="body1" gutterBottom>
              {shipping.province}, {shipping.country}, {shipping.postalCode}
            </Typography>
          </Box>


          <Box sx={{ mt: 2, mb: 3 }}>
            <Typography variant="body1" gutterBottom>
              Selected carrier: {carrier.name}
            </Typography>
  
          </Box>

          <Box sx={{ mt: 2, mb: 3 }} display={"column"}>
            <Box display={"flex"} justifyContent={"space-between"}>
              <Typography variant="body1" gutterBottom>
                Product: {product.Name}
              </Typography>

              <Card>
                <CardMedia
                  component="img"
                  height="70"
                  image={product.ImageURL || "/418eyXxdCsL._SY445_SX342_.jpg"}
                  alt={product.Name}
                />
              </Card>
            </Box>
  

            {product.Sold ? (
              <Typography variant="body1" gutterBottom color="error">
                VENDIDO
              </Typography>
            ) : (
              <Typography variant="body1" gutterBottom color="warning">
                DISPONIBLE
              </Typography>
            )}

            <Typography variant="h6" sx={{ mt: 2 }}>
              Total: {product.Price}
            </Typography>
          </Box>

          <Box sx={{ mt: 3 }}>
            <Button
              onClick={handlePayment}
              variant="contained"
              color="primary"
              fullWidth
              disabled={loading}
            >
              {loading ? (
                <CircularProgress size={24} color="inherit" />
              ) : (
                "Proceed to Payment"
              )}
            </Button>
          </Box>

          {error && (
            <Box sx={{ mt: 2 }}>
              <Alert severity="error">{error}</Alert>
            </Box>
          )}
        </Box>
      </Modal>
    </div>
  );
};
