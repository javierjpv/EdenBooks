import { ApiResponse } from "../../../interfaces/apiResponse";
import { IChatRequest } from "./IChatRequest";
import { IChat } from "./IChat";

export interface IChatApi{
    CreateChat:(chat:IChatRequest)=>Promise<ApiResponse<IChat>>
    GetChatById:(id:number)=>Promise<ApiResponse<IChat>>
    GetChats:()=>Promise<ApiResponse<IChat[]>>
    UpdateChat:(id:number,chat:IChatRequest)=>Promise<ApiResponse<IChat>>
    DeleteChat:(id:number)=>Promise<ApiResponse<IChat>>
} 