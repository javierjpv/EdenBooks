import { ApiResponse } from "../../../interfaces/apiResponse"
import { ICarrier } from "./ICarrier"


export interface ICarrierService{
    CreateCarrier:(carrier:ICarrier)=>Promise<ApiResponse<ICarrier>>
    GetCarrierById:(id:number)=>Promise<ApiResponse<ICarrier>>
    GetCarriers:()=>Promise<ApiResponse<ICarrier[]>>
    UpdateCarrier:(id:number,carrier:ICarrier)=>Promise<ApiResponse<ICarrier>>
    DeleteCarrier:(id:number)=>Promise<ApiResponse<ICarrier>>
} 