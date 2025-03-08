import { ICategory } from "./ICategory";
import { ICategoryRequest } from "./ICategoryRequest";
export interface ICategoryService {
  GetCategories(): Promise<ICategory[]>;
  GetCategoryById(id: number): Promise<ICategory>;
  CreateCategory(category: ICategoryRequest): Promise<void>;
  UpdateCategory(id: number, category: ICategoryRequest): Promise<void>;
  DeleteCategory(id: number): Promise<void>;
}
