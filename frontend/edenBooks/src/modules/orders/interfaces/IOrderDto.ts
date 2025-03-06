export interface IOrderDto{
    ID: number
    CreatedAt: string
    UpdatedAt: string
    DeletedAt: null|string
    State: string
    UserID: number
    AddressID:number 
    CarrierID: number
    TransactionID: number
    Products: null
}