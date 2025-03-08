import { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IChatApi } from "../interfaces/IChatApi";
import { IChatRequest } from "../interfaces/IChatRequest";
import { IChatDto } from "../interfaces/IChatDto";
import { axiosInstance } from "../../../api/axiosInstance";


const BASE_URL = "/chats";

export const chatApi: IChatApi = {
  CreateChat: async ( chat: IChatRequest): Promise<ApiResponse<IChatDto>> => {
    try {
      const response= await axiosInstance.post(BASE_URL, chat);
      const chatDto = response.data;


      return { success: true,data:chatDto};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetChats: async (): Promise<ApiResponse<IChatDto[]>> => {
    try {
      const response = await axiosInstance.get<IChatDto[]>(BASE_URL);
      const chatsDto = response.data;

      return { success: true, data: chatsDto };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetChatById: async (id): Promise<ApiResponse<IChatDto>> => {
    try {
      const response = await axiosInstance.get<IChatDto>(`${BASE_URL}/${id}`);
      const chatDto = response.data;

      return { success: true, data: chatDto };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateChat: async (id, chat: IChatRequest): Promise<ApiResponse<IChatDto>> => {
    try {
      await axiosInstance.put(`${BASE_URL}/${id}`, chat);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteChat: async (id): Promise<ApiResponse<IChatDto>> => {
    try {
      await axiosInstance.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
