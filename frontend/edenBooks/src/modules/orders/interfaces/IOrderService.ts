import { ApiResponse } from "../../../interfaces/apiResponse"
import { IOrder } from "./IOrder"


export interface IOrderService{
    CreateOrder:(order:IOrder)=>Promise<ApiResponse<IOrder>>
    GetOrderById:(id:number)=>Promise<ApiResponse<IOrder>>
    GetOrders:()=>Promise<ApiResponse<IOrder[]>>
    UpdateOrder:(id:number,order:IOrder)=>Promise<ApiResponse<IOrder>>
    DeleteOrder:(id:number)=>Promise<ApiResponse<IOrder>>
} 