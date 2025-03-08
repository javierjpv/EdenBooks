import { ICategoryResponse } from "./ICategoryResponse";
import { ICategoryRequest } from "./ICategoryRequest";
export interface ICategoryApi {
  GetCategories(): Promise<ICategoryResponse[]>;
  GetCategoryById(id: number): Promise<ICategoryResponse>;
  CreateCategory(category: ICategoryRequest): Promise<void>;
  UpdateCategory(id: number, category: ICategoryRequest): Promise<void>;
  DeleteCategory(id: number): Promise<void>;
}
