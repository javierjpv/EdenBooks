import { IMessageDto } from "../../messages/interfaces/IMessageDto";
import { IUser } from "../../users/interfaces/IUser"

export interface IChatDto{
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    Users:IUser[] 
    Messages:null|IMessageDto[]
}