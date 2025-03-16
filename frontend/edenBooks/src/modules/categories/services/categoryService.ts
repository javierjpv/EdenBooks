import { categoryApi } from "../api/categoryApi";
import { ICategoryRequest } from "../interfaces/ICategoryRequest";
import { ICategoryService } from "../interfaces/ICategoryService";

export const CategoryService: ICategoryService = {
  CreateCategory: async (category: ICategoryRequest) =>categoryApi.CreateCategory(category),
  GetCategories: async () => categoryApi.GetCategories(),
  GetCategoryById: async (id: number) => categoryApi.GetCategoryById(id),
  UpdateCategory: async (id: number, category: ICategoryRequest) =>categoryApi.UpdateCategory(id, category),
  DeleteCategory: async (id: number) => categoryApi.DeleteCategory(id),
};
