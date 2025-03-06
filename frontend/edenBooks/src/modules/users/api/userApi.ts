import axios, { AxiosError } from "axios";
import { IUserApi } from "../interfaces/IUserApi";
import { IUser } from "../interfaces/IUser";
import { ApiResponse } from "../../../interfaces/apiResponse";
const BASE_URL = "http://localhost:6969/users";

export const userApi: IUserApi = {
  Login: async (
    email: string,
    password: string
  ): Promise<ApiResponse<IUser>> => {
    const user: IUser = { Email: email, Password: password };
    try {
      const response = await axios.post<IUser>(`${BASE_URL}/login`, user);
      return { data: response.data, success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  Register: async (
    email: string,
    password: string
  ): Promise<ApiResponse<IUser>> => {
    const user: IUser = { Email: email, Password: password };
    try {
      const response = await axios.post<IUser>(`${BASE_URL}/register`, user);
      return { data: response.data, success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },  GetUserById: async (id): Promise<ApiResponse<IUser>> => {
      try {
        const response = await axios.get<IUser>(`${BASE_URL}/${id}`);
        const user = response.data;

  
        return { success: true, data: user };
      } catch (error) {
        const axiosError = error as AxiosError<{ error: string }>;
        return { success: false, error: axiosError.request?.response };
      }
    },
};
