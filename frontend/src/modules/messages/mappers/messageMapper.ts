import { IMessage } from "../interfaces/IMessage";
import { IMessageResponse } from "../interfaces/IMessageResponse";

export const FromMessageResponse=(messageDto:IMessageResponse):IMessage=>{

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