import { ApiResponse } from "../../../interfaces/apiResponse"
import { IUser } from "./IUser"
export interface IUserApi{
    Login(email:string,password:string):Promise<ApiResponse<IUser>> 
    Register(email:string,password:string):Promise<ApiResponse<IUser>> 
    GetUserById:(id:number)=>Promise<ApiResponse<IUser>>
    
}