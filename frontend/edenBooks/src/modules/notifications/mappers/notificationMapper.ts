import { INotification } from "../interfaces/INotification";
import { INotificationDto } from "../interfaces/INotificationDto";





export const NotificationFromDto=(notificationDto:INotificationDto):INotification=>{
const notification:INotification={
    ID:notificationDto.ID,
    Content:notificationDto.Content,
    Seen:notificationDto.Seen,
    UserID:notificationDto.UserID
}
return notification
}