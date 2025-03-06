import { IMessage } from "../interfaces/IMessage";
import { IMessageDto } from "../interfaces/IMessageDto";

export const MessageFromDto=(messageDto:IMessageDto):IMessage=>{

    const message:IMessage={
        ID:messageDto.ID,
        content:messageDto.Content,
        status:messageDto.Status,
        seen:messageDto.Seen,
        receiverID:messageDto.ReceiverID,
        senderID:messageDto.SenderID,
        chatID:messageDto.ChatID,
    }
    return message
}