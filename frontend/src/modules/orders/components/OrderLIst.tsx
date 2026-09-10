import { useEffect, useState } from "react";
import { IOrder } from "../interfaces/IOrder";
import { orderService } from "../services/orderService";
import { OrderItem } from "./OrderItem";
import { Grid2, Typography, Paper, Card, CardContent, Skeleton, Stack } from "@mui/material";
export const OrderLIst = () => {
  const [orders, setorders] = useState<IOrder[]>([]);
  const [error, seterror] = useState<boolean>(false)
  const [loading, setloading] = useState<boolean>(true)

  const fetchOrders = async () => {
    const response = await orderService.GetOrders();
    setloading(false)
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

      {error ? (
        <Typography variant="h6" color="error" align="center">
          Ha ocurrido un error al cargar los pedidos.
        </Typography>
      ) : loading ? (
        <Grid2 container spacing={2}>
          {[...Array(5)].map((_, index) => (
            <Grid2  size={{ xs: 12 }} key={index}>
              <Card variant="outlined" sx={{ minWidth: 275, p: 2, boxShadow: 3 }}>
                <CardContent>
                  <Skeleton variant="text" width="40%" height={28} />
                  <Stack spacing={1} sx={{ mt: 1 }}>
                    {[...Array(9)].map((_, i) => (
                      <Skeleton key={i} variant="text" width="80%" height={20} />
                    ))}
                  </Stack>
                </CardContent>
              </Card>
            </Grid2>
          ))}
        </Grid2>
      ) : orders.length > 0 ? (
        <Grid2 container spacing={2}>
          {orders.map((order) => (
            <Grid2 size={{ xs: 12}} key={order.ID}>
              <OrderItem order={order} />
            </Grid2>
          ))}
        </Grid2>
      ) : (
        <Typography variant="h6" color="textSecondary" align="center">
          No hay pedidos disponibles.
        </Typography>
      )}
    </Paper>
  );
};
