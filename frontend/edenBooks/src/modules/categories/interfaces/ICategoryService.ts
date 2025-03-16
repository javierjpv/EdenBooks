import { ApiResponse } from "../../../interfaces/apiResponse"
import { ICategory } from "./ICategory"
import { ICategoryRequest } from "./ICategoryRequest"


export interface ICategoryService{
    CreateCategory:(category:ICategoryRequest)=>Promise<ApiResponse<ICategory>>
    GetCategoryById:(id:number)=>Promise<ApiResponse<ICategory>>
    GetCategories:()=>Promise<ApiResponse<ICategory[]>>
    UpdateCategory:(id:number,category:ICategoryRequest)=>Promise<ApiResponse<ICategory>>
    DeleteCategory:(id:number)=>Promise<ApiResponse<ICategory>>
} 