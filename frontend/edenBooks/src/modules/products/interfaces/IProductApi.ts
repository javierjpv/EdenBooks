import { IProduct } from "./IProduct";
import { IProductDto } from "./IProductDto";

export interface IProductApi {
    GetProducts():Promise<IProductDto[]>
    GetProductById(id: number):Promise<IProductDto>
    CreateProduct(product: IProduct):Promise<void>
    UpdateProduct(id:number, product:IProduct):Promise<void>
    DeleteProduct(id:number):Promise<void>
}