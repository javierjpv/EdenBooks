import { ApiResponse } from "../../../interfaces/apiResponse"
import { ICarrier } from "./ICarrier"
import { ICarrierRequest } from "./ICarrierRequest"


export interface ICarrierService{
    CreateCarrier:(carrier:ICarrierRequest)=>Promise<ApiResponse<ICarrier>>
    GetCarrierById:(id:number)=>Promise<ApiResponse<ICarrier>>
    GetCarriers:()=>Promise<ApiResponse<ICarrier[]>>
    UpdateCarrier:(id:number,carrier:ICarrierRequest)=>Promise<ApiResponse<ICarrier>>
    DeleteCarrier:(id:number)=>Promise<ApiResponse<ICarrier>>
} 