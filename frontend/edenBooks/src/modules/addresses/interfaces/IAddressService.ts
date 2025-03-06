import { ApiResponse } from "../../../interfaces/apiResponse"
import { IAddress } from "./IAddress"


export interface IAddressService{
    CreateAddress:(address:IAddress)=>Promise<ApiResponse<IAddress>>
    GetAddressById:(id:number)=>Promise<ApiResponse<IAddress>>
    GetAddresss:()=>Promise<ApiResponse<IAddress[]>>
    UpdateAddress:(id:number,address:IAddress)=>Promise<ApiResponse<IAddress>>
    DeleteAddress:(id:number)=>Promise<ApiResponse<IAddress>>
} 