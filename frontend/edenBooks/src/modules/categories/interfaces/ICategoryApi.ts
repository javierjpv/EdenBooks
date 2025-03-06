import { ICategoryDto } from "./ICategoryDto";
import { ICategory } from "./ICategory";
export interface ICategoryApi {
  GetCategories(): Promise<ICategoryDto[]>;
  GetCategoryById(id: number): Promise<ICategoryDto>;
  CreateCategory(category: ICategory): Promise<void>;
  UpdateCategory(id: number, category: ICategory): Promise<void>;
  DeleteCategory(id: number): Promise<void>;
}
