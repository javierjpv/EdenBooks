import { ApiResponse } from "../../../interfaces/apiResponse"
import { IUser } from "./IUser"
import { IUserRequest } from "./IUserRequest"
export interface IUserService{
    Login(email:string,password:string):(Promise<ApiResponse<IUser>> )
    Register(email:string,password:string):Promise<ApiResponse<IUser>> 
    GetUserById:(id:number)=>Promise<ApiResponse<IUser>>
    UpdateUser:(id:number,user:IUserRequest)=>Promise<ApiResponse<IUser>>
    
}