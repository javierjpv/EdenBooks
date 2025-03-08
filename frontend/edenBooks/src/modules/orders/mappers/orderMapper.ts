import { IOrder } from "../interfaces/IOrder";
import { IOrderResponse } from "../interfaces/IOrderResponse";

export const FromOrderResponse=(orderDto:IOrderResponse):IOrder=>{
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