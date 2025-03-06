import { ApiResponse } from "../../../interfaces/apiResponse"
import { IChat } from "./IChat"
import { IChatDto } from "./IChatDto"


export interface IChatService{
    CreateChat:(chat:IChat)=>Promise<ApiResponse<IChatDto>>
    GetChatById:(id:number)=>Promise<ApiResponse<IChatDto>>
    GetChats:()=>Promise<ApiResponse<IChatDto[]>>
    UpdateChat:(id:number,chat:IChat)=>Promise<ApiResponse<IChatDto>>
    DeleteChat:(id:number)=>Promise<ApiResponse<IChatDto>>
} 
