import { IMessage } from "../../messages/interfaces/IMessage";
import { IUser } from "../../users/interfaces/IUser";

export interface IChat{
        ID: number;
        CreatedAt: string;
        UpdatedAt: string;
        Users:IUser[] 
        Messages:null|IMessage[]
}