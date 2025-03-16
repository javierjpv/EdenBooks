import { IProductResponse } from "./IProductResponse";
import { IProductRequest } from "./IProductResquest";

export interface IProductApi {
    GetProducts(params?: URLSearchParams):Promise<IProductResponse[]>
    GetFavorites():Promise<IProductResponse[]>
    GetProductById(id: number):Promise<IProductResponse>
    CreateProduct(product: IProductRequest):Promise<void>
    UpdateProduct(id:number, product:IProductRequest):Promise<void>
    DeleteProduct(id:number):Promise<void>
    AddToFavorite(id:number):Promise<void>
    DeleteFromFavorites(id:number):Promise<void>
}