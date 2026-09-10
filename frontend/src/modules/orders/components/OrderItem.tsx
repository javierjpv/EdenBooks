import { IOrder } from "../interfaces/IOrder"
import { Card, CardContent, Typography, Stack } from "@mui/material";
export const OrderItem = ({ order }: { order: IOrder }) => {
  return (
    <>
    <Card variant="outlined" sx={{ minWidth: 275, p: 2, boxShadow: 3 }}>
      <CardContent>
        <Typography variant="h6" color="primary">
          Pedido #{order.ID}
        </Typography>
        <Stack spacing={1} sx={{ mt: 1 }}>
          <Typography variant="body1">
            <strong>Estado:</strong> {order.state}
          </Typography>
          <Typography variant="body1">
            <strong>ID Dirección:</strong> {order.addressID}
          </Typography>
          <Typography variant="body1">
            <strong>ID Transportista:</strong> {order.carrierID}
          </Typography>
          <Typography variant="body1">
            <strong>ID Usuario:</strong> {order.userID}
          </Typography>
          <Typography variant="body1">
            <strong>ID Transaction:</strong> {order.transactionID}
          </Typography>
        </Stack>
      </CardContent>
    </Card>

    </>
  )
}
