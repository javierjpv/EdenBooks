import { INotification } from "../interfaces/INotification";
import { INotificationService } from "../interfaces/INotificationService";
import { notificationApi } from "../api/notificationApi";

export const notificationService: INotificationService = {
  CreateNotification: async (notification: INotification) =>notificationApi.CreateNotification(notification),
  GetNotifications: async () => notificationApi.GetNotifications(),
  GetNotificationById: async (id: number) => notificationApi.GetNotificationById(id),
  UpdateNotification: async (id: number, notification: INotification) =>notificationApi.UpdateNotification(id, notification),
  DeleteNotification: async (id: number) => notificationApi.DeleteNotification(id),
};
