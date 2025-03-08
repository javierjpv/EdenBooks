import { INotification } from "../interfaces/INotification";
import { INotificationResponse } from "../interfaces/INotificationResponse";





export const FromNotificationResponse=(notificationDto:INotificationResponse):INotification=>{
const notification:INotification={
    ID:notificationDto.ID,
    Content:notificationDto.Content,
    Seen:notificationDto.Seen,
    UserID:notificationDto.UserID
}
return notification
}