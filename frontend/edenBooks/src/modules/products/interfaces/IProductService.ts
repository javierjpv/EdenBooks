import { ApiResponse } from "../../../interfaces/apiResponse"
import { IProduct } from "./IProduct"
import { IProductRequest } from "./IProductResquest"
export interface IProductService{
        CreateProduct:(product:IProductRequest)=>Promise<ApiResponse<IProduct>>
        GetProductById:(id:number)=>Promise<ApiResponse<IProduct>>
        GetProducts:(params?: URLSearchParams)=>Promise<ApiResponse<IProduct[]>>
        UpdateProduct:(id:number,product:IProductRequest)=>Promise<ApiResponse<IProduct>>
        DeleteProduct:(id:number)=>Promise<ApiResponse<IProduct>>


        GetFavorites():Promise<IProduct[]>
        AddToFavorite(id:number):Promise<void>
        DeleteFromFavorite(id:number):Promise<void>
}