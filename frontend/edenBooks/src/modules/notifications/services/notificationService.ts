import { INotificationService } from "../interfaces/INotificationService";
import { notificationApi } from "../api/notificationApi";
import { INotificationRequest } from "../interfaces/INotificationRequest";

export const notificationService: INotificationService = {
  CreateNotification: async (notification: INotificationRequest) =>notificationApi.CreateNotification(notification),
  GetNotifications: async () => notificationApi.GetNotifications(),
  GetNotificationById: async (id: number) => notificationApi.GetNotificationById(id),
  UpdateNotification: async (id: number, notification: INotificationRequest) =>notificationApi.UpdateNotification(id, notification),
  DeleteNotification: async (id: number) => notificationApi.DeleteNotification(id),
};
