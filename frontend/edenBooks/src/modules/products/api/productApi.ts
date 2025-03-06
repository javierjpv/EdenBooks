import { IProductApi } from "../interfaces/IProductApi";
import { IProductDto } from "../interfaces/IProductDto";
import { IProduct } from "../interfaces/IProduct";
import axios from "axios";

const BASE_URL = "http://localhost:6969/products";

export const productApi: IProductApi = {

  GetProducts: async (): Promise<IProductDto[]> => {
    const response = await axios.get<IProductDto[]>(BASE_URL);
    return response.data;
  },
  GetProductById: async (id: number): Promise<IProductDto> => {
    const response = await axios.get<IProductDto>(`${BASE_URL}/${id}`);
    return response.data;
  },
  CreateProduct: async (product: IProduct): Promise<void> => {
    await axios.post(BASE_URL, product);
  },
  UpdateProduct: async (id: number, product: IProduct): Promise<void> => {
    await axios.put(`${BASE_URL}/${id}`, product);
  },
  DeleteProduct: async (id: number): Promise<void> => {
    await axios.delete(`${BASE_URL}/${id}`);
  },

};
