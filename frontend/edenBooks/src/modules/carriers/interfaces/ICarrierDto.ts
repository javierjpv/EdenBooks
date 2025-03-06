 import { IOrderDto } from "../../orders/interfaces/IOrderDto"

export interface ICarrierDto {
    ID:number
    CreatedAt:string
    UpdatedAt: string
    DeletedAt: string|null

    Name:string 
    Contact:string
    Orders:null|IOrderDto[]
  }