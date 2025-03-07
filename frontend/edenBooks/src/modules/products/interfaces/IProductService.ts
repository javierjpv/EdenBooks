import { IProduct } from "./IProduct"
export interface IProductService{
        GetProducts():Promise<IProduct[]>
        GetFavorites():Promise<IProduct[]>
        GetProductById(id: number):Promise<IProduct>
        CreateProduct(product: IProduct):Promise<void>
        UpdateProduct(id:number, product:IProduct):Promise<void>
        DeleteProduct(id:number):Promise<void>
        AddToFavorite(id:number):Promise<void>
        DeleteFromFavorite(id:number):Promise<void>
}