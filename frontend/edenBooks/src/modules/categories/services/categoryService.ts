import { categoryApi } from "../api/categoryApi";
import { ICategoryService } from "../interfaces/ICategoryService";
import { FromCategoryDto } from "../mappers/categoryMapper";
import { ICategory } from "../interfaces/ICategory";
import { ICategoryDto } from "../interfaces/ICategoryDto";
export const categoryService: ICategoryService = {
  GetCategories: async (): Promise<ICategory[]> => {
    try {
      const categoriesDto: ICategoryDto[] = await categoryApi.GetCategories();
      const categories: ICategory[] = categoriesDto.map((categoryDto) =>
        FromCategoryDto(categoryDto)
      );
      return categories;
    } catch (error) {
      console.log("Error al obtener las categorias", error);
      throw new Error("Error al obtener las categorias");
    }
  },
  // GetCategoriesDto: async (): Promise<ICategoryDto[]> => {
  //   try {
  //     const categoriesDto: ICategoryDto[] = await categoryApi.GetCategories();
  //     return categoriesDto;
  //   } catch (error) {
  //     console.log("Error al obtener las categorias", error);
  //     throw new Error("Error al obtener las categorias");
  //   }
  // },
  CreateCategory: async (category: ICategory): Promise<void> => {
    try {
      await categoryApi.CreateCategory(category);
    } catch (error) {
      console.log("Error al crear la categoria", error);
      throw new Error("Error al crear la categoria");
    }
  },
  GetCategoryById: async (id: number): Promise<ICategory> => {
    try {
      const categoryDto: ICategoryDto = await categoryApi.GetCategoryById(id);
      const category: ICategory = FromCategoryDto(categoryDto);
      return category;
    } catch (error) {
      console.log("Error al obtener categoria por id", error);
      throw new Error("Error al obtener categoria por id");
    }
  },
  UpdateCategory: async (id: number, category: ICategory): Promise<void> => {
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
