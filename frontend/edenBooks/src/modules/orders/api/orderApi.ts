import { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IOrderApi } from "../interfaces/IOrderApi";
import { IOrder } from "../interfaces/IOrder";
import { FromOrderResponse} from "../mappers/orderMapper";
import { IOrderResponse } from "../interfaces/IOrderResponse";
import { axiosInstance } from "../../../api/axiosInstance";
import { IOrderRequest } from "../interfaces/IOrderRequest";


const BASE_URL = "/orders";


export const orderApi: IOrderApi = {
  CreateOrder: async ( order: IOrderRequest): Promise<ApiResponse<IOrder>> => {
    try {
      const response= await axiosInstance.post(BASE_URL, order);
      const orderDto = response.data;
      const createdOrders: IOrder = FromOrderResponse(orderDto);

      return { success: true,data:createdOrders};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetOrders: async (): Promise<ApiResponse<IOrder[]>> => {

    try {
      const response = await axiosInstance.get<IOrderResponse[]>(BASE_URL);
      const ordersDto = response.data;
      const orders: IOrder[] = ordersDto.map((orderDto) =>
        FromOrderResponse(orderDto)
      );

      return { success: true, data: orders };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetOrderById: async (id): Promise<ApiResponse<IOrder>> => {
    try {
      const response = await axiosInstance.get<IOrderResponse>(`${BASE_URL}/${id}`);
      const orderDto = response.data;

      const order: IOrder = FromOrderResponse(orderDto);

      return { success: true, data: order };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateOrder: async (id, order: IOrderRequest): Promise<ApiResponse<IOrder>> => {
    try {
      await axiosInstance.put(`${BASE_URL}/${id}`, order);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteOrder: async (id): Promise<ApiResponse<IOrder>> => {
    try {
      await axiosInstance.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
