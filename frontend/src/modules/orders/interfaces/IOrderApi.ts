import { ApiResponse } from "../../../interfaces/apiResponse";
import { IOrder } from "./IOrder";
import { IOrderRequest } from "./IOrderRequest";

export interface IOrderApi{
    CreateOrder:(order:IOrderRequest)=>Promise<ApiResponse<IOrder>>
    GetOrderById:(id:number)=>Promise<ApiResponse<IOrder>>
    GetOrders:()=>Promise<ApiResponse<IOrder[]>>
    UpdateOrder:(id:number,order:IOrderRequest)=>Promise<ApiResponse<IOrder>>
    DeleteOrder:(id:number)=>Promise<ApiResponse<IOrder>>
} 