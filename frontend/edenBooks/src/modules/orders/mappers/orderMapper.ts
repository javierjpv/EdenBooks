import { IOrder } from "../interfaces/IOrder";
import { IOrderDto } from "../interfaces/IOrderDto";

export const OrderFromDto=(orderDto:IOrderDto):IOrder=>{
const order:IOrder={
ID:orderDto.ID,
state:orderDto.State,
addressID:orderDto.AddressID,
carrierID:orderDto.CarrierID,
transactionID:orderDto.TransactionID,
userID:orderDto.UserID
}
return order
}