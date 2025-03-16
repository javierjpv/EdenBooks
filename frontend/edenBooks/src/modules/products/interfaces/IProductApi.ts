import { ApiResponse } from "../../../interfaces/apiResponse";
import { IProduct } from "./IProduct";
import { IProductResponse } from "./IProductResponse";
import { IProductRequest } from "./IProductResquest";

export interface IProductApi {
    CreateProduct:(product:IProductRequest)=>Promise<ApiResponse<IProduct>>
    GetProductById:(id:number)=>Promise<ApiResponse<IProduct>>
    GetProducts:(params?: URLSearchParams)=>Promise<ApiResponse<IProduct[]>>
    UpdateProduct:(id:number,product:IProductRequest)=>Promise<ApiResponse<IProduct>>
    DeleteProduct:(id:number)=>Promise<ApiResponse<IProduct>>

    
    GetFavorites():Promise<IProductResponse[]>
    AddToFavorite(id:number):Promise<void>
    DeleteFromFavorites(id:number):Promise<void>
}