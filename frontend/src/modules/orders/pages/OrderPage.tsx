import { OrderLIst } from "../components/OrderLIst"
import { Container } from "@mui/material"

export const OrderPage = () => {
  return (
    <>
    <Container maxWidth="md" sx={{ marginTop: 22 }}>   
    <h1>OrderPage</h1>
    <OrderLIst/>
    </Container>
    </>
  )
}
