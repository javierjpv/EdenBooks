import { IProduct } from "./IProduct"
import { IProductRequest } from "./IProductResquest"
export interface IProductService{
        GetProducts(params?: URLSearchParams):Promise<IProduct[]>
        GetFavorites():Promise<IProduct[]>
        GetProductById(id: number):Promise<IProduct>
        CreateProduct(product: IProductRequest):Promise<void>
        UpdateProduct(id:number, product:IProductRequest):Promise<void>
        DeleteProduct(id:number):Promise<void>
        AddToFavorite(id:number):Promise<void>
        DeleteFromFavorite(id:number):Promise<void>
}