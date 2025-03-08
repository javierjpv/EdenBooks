 import { IOrderResponse } from "../../orders/interfaces/IOrderResponse"

export interface ICarrierResponse {
    ID:number
    CreatedAt:string
    UpdatedAt: string
    DeletedAt: string|null

    Name:string 
    Contact:string
    Orders:null|IOrderResponse[]
  }