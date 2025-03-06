import { ApiResponse } from "../../../interfaces/apiResponse"
import { IMessage } from "./IMessage"


export interface IMessageService{
    CreateMessage:(message:IMessage)=>Promise<ApiResponse<IMessage>>
    GetMessageById:(id:number)=>Promise<ApiResponse<IMessage>>
    GetMessages:()=>Promise<ApiResponse<IMessage[]>>
    UpdateMessage:(id:number,message:IMessage)=>Promise<ApiResponse<IMessage>>
    DeleteMessage:(id:number)=>Promise<ApiResponse<IMessage>>
} 