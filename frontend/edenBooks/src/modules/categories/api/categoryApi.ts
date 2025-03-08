import axios from "axios";
import { ICategoryApi } from "../interfaces/ICategoryApi";
import { ICategoryResponse } from "../interfaces/ICategoryResponse";
import { ICategoryRequest } from "../interfaces/ICategoryRequest";
const BASE_URL = "http://localhost:6969/categories";
  

  export const categoryApi:ICategoryApi={
    GetCategories:async():Promise<ICategoryResponse[]>=>{
        const response=await axios.get<ICategoryResponse[]>(BASE_URL)
        return response.data
    },
    CreateCategory:async(category:ICategoryRequest):Promise<void>=>{
        await axios.post(BASE_URL,category)
    },
    GetCategoryById:async(id: number): Promise<ICategoryResponse>=>{
        const response=await axios.get(`${BASE_URL}/${id}`)
        return response.data
    },
    UpdateCategory:async(id: number, category: ICategoryRequest): Promise<void>=>{
        await axios.put(`${BASE_URL}/${id}`,category)
    }
    ,
    DeleteCategory:async(id: number): Promise<void>=>{
        await axios.delete(`${BASE_URL}/${id}`)
    }
  } 
  
  

