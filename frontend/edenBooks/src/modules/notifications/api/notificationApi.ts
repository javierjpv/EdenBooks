import  { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { INotificationApi } from "../interfaces/INotificationApi";
import { INotification } from "../interfaces/INotification";
import { NotificationFromDto } from "../mappers/notificationMapper";
import { INotificationDto } from "../interfaces/INotificationDto";
import { axiosInstance } from "../../../api/axiosInstance";
import { INotificationRequest } from "../interfaces/INotificationRequest";

const BASE_URL = "/notifications";

export const notificationApi: INotificationApi = {
  CreateNotification: async ( notification: INotificationRequest): Promise<ApiResponse<INotification>> => {
    try {
      const response= await axiosInstance.post(BASE_URL, notification);
      const notificationDto = response.data;
      const createdNotifications: INotification = NotificationFromDto(notificationDto);

      return { success: true,data:createdNotifications};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetNotifications: async (): Promise<ApiResponse<INotification[]>> => {
    try {
      const response = await axiosInstance.get<INotificationDto[]>(BASE_URL);
      const notificationsDto = response.data;
      const notifications: INotification[] = notificationsDto.map((notificationDto) =>
        NotificationFromDto(notificationDto)
      );

      return { success: true, data: notifications };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetNotificationById: async (id): Promise<ApiResponse<INotification>> => {
    try {
      const response = await axiosInstance.get<INotificationDto>(`${BASE_URL}/${id}`);
      const notificationDto = response.data;

      const notification: INotification = NotificationFromDto(notificationDto);

      return { success: true, data: notification };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateNotification: async (id, notification: INotificationRequest): Promise<ApiResponse<INotification>> => {
    try {
      await axiosInstance.put(`${BASE_URL}/${id}`, notification);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteNotification: async (id): Promise<ApiResponse<INotification>> => {
    try {
      await axiosInstance.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
