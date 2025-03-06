import axios from "axios";
import { ICategoryApi } from "../interfaces/ICategoryApi";
import { ICategoryDto } from "../interfaces/ICategoryDto";
import { ICategory } from "../interfaces/ICategory";
const BASE_URL = "http://localhost:6969/categories";
  

  export const categoryApi:ICategoryApi={
    GetCategories:async():Promise<ICategoryDto[]>=>{
        const response=await axios.get<ICategoryDto[]>(BASE_URL)
        return response.data
    },
    CreateCategory:async(category:ICategory):Promise<void>=>{
        await axios.post(BASE_URL,category)
    },
    GetCategoryById:async(id: number): Promise<ICategoryDto>=>{
        const response=await axios.get(`${BASE_URL}/${id}`)
        return response.data
    },
    UpdateCategory:async(id: number, category: ICategory): Promise<void>=>{
        await axios.put(`${BASE_URL}/${id}`,category)
    }
    ,
    DeleteCategory:async(id: number): Promise<void>=>{
        await axios.delete(`${BASE_URL}/${id}`)
    }
  } 
  
  

