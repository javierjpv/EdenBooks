import { ApiResponse } from "../../../interfaces/apiResponse";
import { INotification } from "./INotification";

export interface INotificationApi{
    CreateNotification:(notification:INotification)=>Promise<ApiResponse<INotification>>
    GetNotificationById:(id:number)=>Promise<ApiResponse<INotification>>
    GetNotifications:()=>Promise<ApiResponse<INotification[]>>
    UpdateNotification:(id:number,notification:INotification)=>Promise<ApiResponse<INotification>>
    DeleteNotification:(id:number)=>Promise<ApiResponse<INotification>>
} 