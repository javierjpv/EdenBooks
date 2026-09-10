import { IProductApi } from "../interfaces/IProductApi";
import { IProductResponse } from "../interfaces/IProductResponse";
import { AxiosError } from "axios";
import { axiosInstance } from "../../../api/axiosInstance";
import { IProductRequest } from "../interfaces/IProductResquest";
import { IProduct } from "../interfaces/IProduct";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { FromProductResponse } from "../mappers/productMapper";

const BASE_URL = "http://localhost:6969/products";

export const productApi: IProductApi = {

  CreateProduct: async ( product: IProductRequest): Promise<ApiResponse<IProduct>> => {
    try {
      const response= await axiosInstance.post(BASE_URL, product);
      const productDto = response.data;
      const createdProduct: IProduct = FromProductResponse(productDto);

      return { success: true,data:createdProduct};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetProducts: async (params?: URLSearchParams): Promise<ApiResponse<IProduct[]>> => {
    try {
      const response = await axiosInstance.get<IProductResponse[]>(BASE_URL, {
        params: params ? Object.fromEntries(params) : {}, // Convierte los searchParams en objeto
      });
      const productsDto = response.data;
      const products: IProduct[] = productsDto.map((productDto) =>
        FromProductResponse(productDto)
      );

      return { success: true, data: products };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetProductById: async (id): Promise<ApiResponse<IProduct>> => {
    try {
      const response = await axiosInstance.get<IProductResponse>(`${BASE_URL}/${id}`);
      const productDto = response.data;

      const product: IProduct = FromProductResponse(productDto);

      return { success: true, data: product };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateProduct: async (id, product: IProductRequest): Promise<ApiResponse<IProduct>> => {
    try {
      await axiosInstance.put(`${BASE_URL}/${id}`, product);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteProduct: async (id): Promise<ApiResponse<IProduct>> => {
    try {
      await axiosInstance.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },




  GetFavorites: async (): Promise<IProductResponse[]> => {
    const response = await axiosInstance.get<IProductResponse[]>(`${BASE_URL}/favorites`);
    return response.data;
  },
  AddToFavorite: async (id: number): Promise<void> => {
    await axiosInstance.post(`${BASE_URL}/${id}/favorite`);
  },
  DeleteFromFavorites: async (id: number): Promise<void> => {
    await axiosInstance.delete(`${BASE_URL}/${id}/favorite`);
  },  

};
