import { ICategoryDto } from "./ICategoryDto";
import { ICategoryRequest } from "./ICategoryRequest";
export interface ICategoryApi {
  GetCategories(): Promise<ICategoryDto[]>;
  GetCategoryById(id: number): Promise<ICategoryDto>;
  CreateCategory(category: ICategoryRequest): Promise<void>;
  UpdateCategory(id: number, category: ICategoryRequest): Promise<void>;
  DeleteCategory(id: number): Promise<void>;
}
