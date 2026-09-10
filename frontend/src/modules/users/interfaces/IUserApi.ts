import { ApiResponse } from "../../../interfaces/apiResponse"
import { IUser } from "./IUser"
import { IUserRequest } from "./IUserRequest"
import { IUserResponse } from "./IUserResponse"
export interface IUserApi{
    Login(email:string,password:string):Promise<ApiResponse<IUserResponse>> 
    Register(email:string,password:string):Promise<ApiResponse<IUserResponse>> 
    GetUserById:(id:number)=>Promise<ApiResponse<IUserResponse>>
    UpdateUser:(id:number,user:IUserRequest)=>Promise<ApiResponse<IUser>>
    
}