export interface IProductDto{
    ID:number
    CreatedAt:string
    UpdatedAt: string
    DeletedAt: string|null
    Name: string
    Description: string
    Price: number
    OrderID: number|null
    CategoryID: number
    UserID: number
    ImageURL: string
    Sold:boolean
    Reviews: null
    is_favorite: boolean
}

