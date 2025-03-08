import axios, { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IMessageApi } from "../interfaces/IMessageApi";
import { IMessage } from "../interfaces/IMessage";
import { FromMessageResponse} from "../mappers/messageMapper";
import { IMessageResponse } from "../interfaces/IMessageResponse";
import { IMessageRequest } from "../interfaces/IMessageRequest";

const BASE_URL = "http://localhost:6969/messages";

export const messageApi: IMessageApi = {
  CreateMessage: async ( message: IMessageRequest): Promise<ApiResponse<IMessage>> => {
    try {
      const response= await axios.post(BASE_URL, message);
      const messageDto = response.data;
      const createdMessages: IMessage = FromMessageResponse(messageDto);

      return { success: true,data:createdMessages};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetMessages: async (): Promise<ApiResponse<IMessage[]>> => {
    try {
      const response = await axios.get<IMessageResponse[]>(BASE_URL);
      const messagesDto = response.data;
      const messages: IMessage[] = messagesDto.map((messageDto) =>
        FromMessageResponse(messageDto)
      );

      return { success: true, data: messages };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetMessageById: async (id): Promise<ApiResponse<IMessage>> => {
    try {
      const response = await axios.get<IMessageResponse>(`${BASE_URL}/${id}`);
      const messageDto = response.data;

      const message: IMessage = FromMessageResponse(messageDto);

      return { success: true, data: message };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateMessage: async (id, message: IMessageRequest): Promise<ApiResponse<IMessage>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, message);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteMessage: async (id): Promise<ApiResponse<IMessage>> => {
    try {
      await axios.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
