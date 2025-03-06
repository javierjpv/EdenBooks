import { IProduct } from "./IProduct"
export interface IProductService{
        GetProducts():Promise<IProduct[]>
        GetProductById(id: number):Promise<IProduct>
        CreateProduct(product: IProduct):Promise<void>
        UpdateProduct(id:number, product:IProduct):Promise<void>
        DeleteProduct(id:number):Promise<void>
}