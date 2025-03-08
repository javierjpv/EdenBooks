import { IChat } from "../interfaces/IChat";
import { IChatDto } from "../interfaces/IChatDto";

export const ChatFromDto=(chatDto:IChatDto):IChat=>{

    const chat:IChat={
        ID:chatDto.ID,
       CreatedAt:chatDto.CreatedAt,
       DeletedAt:chatDto.DeletedAt,
       UpdatedAt:chatDto.UpdatedAt,
       Messages:chatDto.Messages,
       Users:chatDto.Users
    }
    return chat
}