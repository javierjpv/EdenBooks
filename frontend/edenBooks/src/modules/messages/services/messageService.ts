import { IMessage } from "../interfaces/IMessage";
import { IMessageService } from "../interfaces/IMessageService";
import { messageApi } from "../api/messageApi";

export const messageService: IMessageService = {
  CreateMessage: async (message: IMessage) =>messageApi.CreateMessage(message),
  GetMessages: async () => messageApi.GetMessages(),
  GetMessageById: async (id: number) => messageApi.GetMessageById(id),
  UpdateMessage: async (id: number, message: IMessage) =>messageApi.UpdateMessage(id, message),
  DeleteMessage: async (id: number) => messageApi.DeleteMessage(id),
};
