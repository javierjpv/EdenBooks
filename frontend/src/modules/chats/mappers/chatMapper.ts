import { IMessage } from "../../messages/interfaces/IMessage";
import { FromMessageResponse } from "../../messages/mappers/messageMapper";
import { IChat } from "../interfaces/IChat";
import { IChatResponse } from "../interfaces/IChatResponse";

export const FromChatResponse=(chatDto:IChatResponse):IChat=>{
    const messages: IMessage[] = chatDto.Messages
    ? chatDto.Messages.map((messageDto) => FromMessageResponse(messageDto))
    : []; //

    const chat:IChat={
        ID:chatDto.ID,
       CreatedAt:chatDto.CreatedAt,
       UpdatedAt:chatDto.UpdatedAt,
       Messages:messages,
       Users:chatDto.Users
    }
    return chat
}