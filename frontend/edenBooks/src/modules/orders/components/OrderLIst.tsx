import { useEffect, useState } from "react";
import { IOrder } from "../interfaces/IOrder";
import { orderService } from "../services/orderService";
import { OrderItem } from "./OrderItem";
import { Grid2, Typography, Paper } from "@mui/material";
export const OrderLIst = () => {
  const [orders, setorders] = useState<IOrder[]>([]);
  const [error, seterror] = useState<boolean>(false)
  const fetchOrders = async () => {
    console.log("Antes de hacer la solicitud a la API");
    console.log("Enviando solicitud a la API...");
    const response = await orderService.GetOrders();
    console.log("Después de hacer la solicitud a la API");
    if (response.success && response.data) {
      setorders(response.data);
      console.log(response.data)
    }
    else{
      seterror(true)
  
    }
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  return (
    
    <Paper elevation={3} sx={{ p: 3, maxWidth: 800, mx: "auto", mt: 4 }}>
      <Typography variant="h4" color="primary" gutterBottom>
        Lista de Pedidos
      </Typography>
      {!error?(<>
        {orders.length > 0 ? (
        <Grid2 container spacing={2}>
          {orders.map((order) => (
            <Grid2 size={{ xs: 12 }} key={order.ID}>
              <OrderItem order={order} />
            </Grid2>
          ))}
        </Grid2>
      ) : (
        <Typography variant="h6" color="textSecondary" align="center">
          No hay pedidos disponibles.
        </Typography>
      )}
      </>
      ):(
        <p>Ha ocurrido un error</p>
      )}

    </Paper>
    
  );
};
