import { IChat } from "../interfaces/IChat";
import { IChatService } from "../interfaces/IChatService";
import { chatApi } from "../api/chatApi";


export const chatService: IChatService = {
  CreateChat: async (chat: IChat) =>chatApi.CreateChat(chat),
  GetChats: async () => chatApi.GetChats(),
  GetChatById: async (id: number) => chatApi.GetChatById(id),
  UpdateChat: async (id: number, chat: IChat) =>chatApi.UpdateChat(id, chat),
  DeleteChat: async (id: number) => chatApi.DeleteChat(id),
};
