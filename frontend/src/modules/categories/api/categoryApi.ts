import axios, { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { ICategoryApi } from "../interfaces/ICategoryApi";
import { ICategory } from "../interfaces/ICategory";
import { FromCategoryResponse } from "../mappers/categoryMapper";
import { ICategoryResponse } from "../interfaces/ICategoryResponse";
import { ICategoryRequest } from "../interfaces/ICategoryRequest";

const BASE_URL = "http://localhost:6969/categories";

export const categoryApi: ICategoryApi = {
  CreateCategory: async ( category: ICategoryRequest): Promise<ApiResponse<ICategory>> => {
    try {
      const response= await axios.post(BASE_URL, category);
      const categoryDto = response.data;
      const createdCategory: ICategory = FromCategoryResponse(categoryDto);

      return { success: true,data:createdCategory};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetCategories: async (): Promise<ApiResponse<ICategory[]>> => {
    try {
      const response = await axios.get<ICategoryResponse[]>(BASE_URL);
      const categorysDto = response.data;
      const categories: ICategory[] = categorysDto.map((categoryDto) =>
        FromCategoryResponse(categoryDto)
      );

      return { success: true, data: categories };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetCategoryById: async (id): Promise<ApiResponse<ICategory>> => {
    try {
      const response = await axios.get<ICategoryResponse>(`${BASE_URL}/${id}`);
      const categoryDto = response.data;

      const category: ICategory = FromCategoryResponse(categoryDto);

      return { success: true, data: category };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateCategory: async (id, category: ICategoryRequest): Promise<ApiResponse<ICategory>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, category);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteCategory: async (id): Promise<ApiResponse<ICategory>> => {
    try {
      await axios.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
