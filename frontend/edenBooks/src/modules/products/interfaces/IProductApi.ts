import { IProductDto } from "./IProductDto";
import { IProductRequest } from "./IProductResquest";

export interface IProductApi {
    GetProducts():Promise<IProductDto[]>
    GetFavorites():Promise<IProductDto[]>
    GetProductById(id: number):Promise<IProductDto>
    CreateProduct(product: IProductRequest):Promise<void>
    UpdateProduct(id:number, product:IProductRequest):Promise<void>
    DeleteProduct(id:number):Promise<void>
    AddToFavorite(id:number):Promise<void>
    DeleteFromFavorites(id:number):Promise<void>
}