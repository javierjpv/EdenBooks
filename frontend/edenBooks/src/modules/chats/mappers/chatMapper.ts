import { IMessage } from "../../messages/interfaces/IMessage";
import { MessageFromDto } from "../../messages/mappers/messageMapper";
import { IChat } from "../interfaces/IChat";
import { IChatDto } from "../interfaces/IChatDto";

export const ChatFromDto=(chatDto:IChatDto):IChat=>{
    const messages: IMessage[] = chatDto.Messages
    ? chatDto.Messages.map((messageDto) => MessageFromDto(messageDto))
    : []; //

    const chat:IChat={
        ID:chatDto.ID,
       CreatedAt:chatDto.CreatedAt,
       DeletedAt:chatDto.DeletedAt,
       UpdatedAt:chatDto.UpdatedAt,
       Messages:messages,
       Users:chatDto.Users
    }
    return chat
}