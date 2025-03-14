import { categoryApi } from "../api/categoryApi";
import { ICategoryService } from "../interfaces/ICategoryService";
import { FromCategoryResponse } from "../mappers/categoryMapper";
import { ICategory } from "../interfaces/ICategory";
import { ICategoryResponse } from "../interfaces/ICategoryResponse";
import { ICategoryRequest } from "../interfaces/ICategoryRequest";
export const categoryService: ICategoryService = {
  GetCategories: async (): Promise<ICategory[]> => {
    try {
      const categoriesDto: ICategoryResponse[] = await categoryApi.GetCategories();
      const categories: ICategory[] = categoriesDto.map((categoryDto) =>
        FromCategoryResponse(categoryDto)
      );
      return categories;
    } catch (error) {
      console.log("Error al obtener las categorias", error);
      throw new Error("Error al obtener las categorias");
    }
  },

  CreateCategory: async (category: ICategoryRequest): Promise<void> => {
    try {
      await categoryApi.CreateCategory(category);
    } catch (error) {
      console.log("Error al crear la categoria", error);
      throw new Error("Error al crear la categoria");
    }
  },
  GetCategoryById: async (id: number): Promise<ICategory> => {
    try {
      const categoryDto: ICategoryResponse = await categoryApi.GetCategoryById(id);
      const category: ICategory = FromCategoryResponse(categoryDto);
      return category;
    } catch (error) {
      console.log("Error al obtener categoria por id", error);
      throw new Error("Error al obtener categoria por id");
    }
  },
  UpdateCategory: async (id: number, category: ICategoryRequest): Promise<void> => {
    try {
      await categoryApi.UpdateCategory(id, category);
    } catch (error) {
      console.log("Error al actualizar la categoria", error);
      throw new Error("Error al actualizar la categoria");
    }
  },
  DeleteCategory: async (id: number): Promise<void> => {
    try {
      await categoryApi.DeleteCategory(id);
    } catch (error) {
      console.log("Error al eliminar la categoria", error);
      throw new Error("Error al eliminar la categoria");
    }
  },
};
