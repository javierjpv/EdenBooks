import { IMessage } from "../interfaces/IMessage";
import { IMessageDto } from "../interfaces/IMessageDto";

export const MessageFromDto=(messageDto:IMessageDto):IMessage=>{

    const message:IMessage={
        ID:messageDto.ID,
        Content:messageDto.Content,
        Status:messageDto.Status,
        Seen:messageDto.Seen,
        ReceiverID:messageDto.ReceiverID,
        SenderID:messageDto.SenderID,
        ChatID:messageDto.ChatID,
    }
    return message
}