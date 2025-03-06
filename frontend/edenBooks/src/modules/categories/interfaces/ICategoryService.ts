import { ICategory } from "./ICategory";
export interface ICategoryService {
  GetCategories(): Promise<ICategory[]>;
  GetCategoryById(id: number): Promise<ICategory>;
  CreateCategory(category: ICategory): Promise<void>;
  UpdateCategory(id: number, category: ICategory): Promise<void>;
  DeleteCategory(id: number): Promise<void>;
}
