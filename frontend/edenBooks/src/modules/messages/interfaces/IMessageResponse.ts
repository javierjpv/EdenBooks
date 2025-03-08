export interface IMessageResponse{
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: null | string;
    Content:string  
	Seen:boolean       
	Status:string    
	ChatID:number       
	SenderID:number  
    ReceiverID:number
}
