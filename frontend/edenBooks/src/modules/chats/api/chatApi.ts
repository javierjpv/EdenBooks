import { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IChatApi } from "../interfaces/IChatApi";
import { IChatRequest } from "../interfaces/IChatRequest";
import { axiosInstance } from "../../../api/axiosInstance";
import { IChat } from "../interfaces/IChat";
import { FromChatResponse } from "../mappers/chatMapper";
import { IChatResponse } from "../interfaces/IChatResponse";


const BASE_URL = "/chats";

export const chatApi: IChatApi = {
  CreateChat: async ( chat: IChatRequest): Promise<ApiResponse<IChat>> => {
    try {
      const response= await axiosInstance.post(BASE_URL, chat);
      const chatDto = response.data;
      const Createdchat=FromChatResponse(chatDto)

      return { success: true,data:Createdchat};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetChats: async (): Promise<ApiResponse<IChat[]>> => {
    try {
      const response = await axiosInstance.get<IChatResponse[]>(BASE_URL);
      const chatsDto = response.data;
      const chats: IChat[] = chatsDto.map((chatDto) =>
        FromChatResponse(chatDto)
      );

      return { success: true, data: chats };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetChatById: async (id): Promise<ApiResponse<IChat>> => {
    try {
      const response = await axiosInstance.get<IChatResponse>(`${BASE_URL}/${id}`);
      const chatDto = response.data;
      const chat=FromChatResponse(chatDto)
      return { success: true, data: chat };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateChat: async (id, chat: IChatRequest): Promise<ApiResponse<IChat>> => {
    try {
      await axiosInstance.put(`${BASE_URL}/${id}`, chat);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteChat: async (id): Promise<ApiResponse<IChat>> => {
    try {
      await axiosInstance.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
