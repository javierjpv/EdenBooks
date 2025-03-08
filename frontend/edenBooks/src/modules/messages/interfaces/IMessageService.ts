import { ApiResponse } from "../../../interfaces/apiResponse"
import { IMessage } from "./IMessage"
import { IMessageRequest } from "./IMessageRequest"


export interface IMessageService{
    CreateMessage:(message:IMessageRequest)=>Promise<ApiResponse<IMessage>>
    GetMessageById:(id:number)=>Promise<ApiResponse<IMessage>>
    GetMessages:()=>Promise<ApiResponse<IMessage[]>>
    UpdateMessage:(id:number,message:IMessageRequest)=>Promise<ApiResponse<IMessage>>
    DeleteMessage:(id:number)=>Promise<ApiResponse<IMessage>>
} 