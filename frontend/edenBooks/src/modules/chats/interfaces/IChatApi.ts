import { ApiResponse } from "../../../interfaces/apiResponse";
import { IChatRequest } from "./IChatRequest";
import { IChatDto } from "./IChatDto";

export interface IChatApi{
    CreateChat:(chat:IChatRequest)=>Promise<ApiResponse<IChatDto>>
    GetChatById:(id:number)=>Promise<ApiResponse<IChatDto>>
    GetChats:()=>Promise<ApiResponse<IChatDto[]>>
    UpdateChat:(id:number,chat:IChatRequest)=>Promise<ApiResponse<IChatDto>>
    DeleteChat:(id:number)=>Promise<ApiResponse<IChatDto>>
} 