import { IMessageService } from "../interfaces/IMessageService";
import { messageApi } from "../api/messageApi";
import { IMessageRequest } from "../interfaces/IMessageRequest";

export const messageService: IMessageService = {
  CreateMessage: async (message: IMessageRequest) =>messageApi.CreateMessage(message),
  GetMessages: async () => messageApi.GetMessages(),
  GetMessageById: async (id: number) => messageApi.GetMessageById(id),
  UpdateMessage: async (id: number, message: IMessageRequest) =>messageApi.UpdateMessage(id, message),
  DeleteMessage: async (id: number) => messageApi.DeleteMessage(id),
};
