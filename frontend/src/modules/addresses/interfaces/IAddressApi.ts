import { ApiResponse } from "../../../interfaces/apiResponse";
import { IAddress } from "./IAddress";
import { IAddressRequest } from "./IAddressRequest";

export interface IAddressApi{
    CreateAddress:(address:IAddressRequest)=>Promise<ApiResponse<IAddress>>
    GetAddressById:(id:number)=>Promise<ApiResponse<IAddress>>
    GetAddresss:()=>Promise<ApiResponse<IAddress[]>>
    UpdateAddress:(id:number,address:IAddressRequest)=>Promise<ApiResponse<IAddress>>
    DeleteAddress:(id:number)=>Promise<ApiResponse<IAddress>>
} 