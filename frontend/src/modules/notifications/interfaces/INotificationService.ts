import { ApiResponse } from "../../../interfaces/apiResponse"
import { INotification } from "./INotification"
import { INotificationRequest } from "./INotificationRequest"


export interface INotificationService{
    CreateNotification:(notification:INotificationRequest)=>Promise<ApiResponse<INotification>>
    GetNotificationById:(id:number)=>Promise<ApiResponse<INotification>>
    GetNotifications:()=>Promise<ApiResponse<INotification[]>>
    UpdateNotification:(id:number,notification:INotificationRequest)=>Promise<ApiResponse<INotification>>
    DeleteNotification:(id:number)=>Promise<ApiResponse<INotification>>
} 