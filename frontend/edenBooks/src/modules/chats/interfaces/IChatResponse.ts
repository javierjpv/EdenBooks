import { IMessageResponse } from "../../messages/interfaces/IMessageResponse";
import { IUser } from "../../users/interfaces/IUser"

export interface IChatResponse{
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    Users:IUser[] 
    Messages:null|IMessageResponse[]
}