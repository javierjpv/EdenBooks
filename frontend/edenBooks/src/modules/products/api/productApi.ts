import { IProductApi } from "../interfaces/IProductApi";
import { IProductResponse } from "../interfaces/IProductDto";
import axios from "axios";
import { axiosInstance } from "../../../api/axiosInstance";
import { IProductRequest } from "../interfaces/IProductResquest";

const BASE_URL = "http://localhost:6969/products";

export const productApi: IProductApi = {

  GetProducts: async (): Promise<IProductResponse[]> => {
    const response = await axiosInstance.get<IProductResponse[]>(BASE_URL);
    return response.data;
  },
  GetFavorites: async (): Promise<IProductResponse[]> => {
    const response = await axiosInstance.get<IProductResponse[]>(`${BASE_URL}/favorites`);
    return response.data;
  },
  GetProductById: async (id: number): Promise<IProductResponse> => {
    const response = await axiosInstance.get<IProductResponse>(`${BASE_URL}/${id}`);
    return response.data;
  },
  CreateProduct: async (product: IProductRequest): Promise<void> => {
    await axios.post(BASE_URL, product);
  },
  UpdateProduct: async (id: number, product: IProductRequest): Promise<void> => {
    await axios.put(`${BASE_URL}/${id}`, product);
  },
  DeleteProduct: async (id: number): Promise<void> => {
    await axios.delete(`${BASE_URL}/${id}`);
  },
  AddToFavorite: async (id: number): Promise<void> => {
    await axiosInstance.post(`${BASE_URL}/${id}/favorite`);
  },
  DeleteFromFavorites: async (id: number): Promise<void> => {
    await axiosInstance.delete(`${BASE_URL}/${id}/favorite`);
  },  

};
