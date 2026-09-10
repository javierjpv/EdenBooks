import axios, { AxiosError } from "axios";
import { IUserApi } from "../interfaces/IUserApi";
import { IUser } from "../interfaces/IUser";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IUserResponse } from "../interfaces/IUserResponse";
import { IUserRequest } from "../interfaces/IUserRequest";
import { FromUserResponse } from "../mappers/userMapper";
const BASE_URL = "http://localhost:6969/users";

export const userApi: IUserApi = {
  Login: async (
    email: string,
    password: string
  ): Promise<ApiResponse<IUser>> => {
    const user: IUserRequest = { Email: email, Password: password };
    try {
      const response = await axios.post(`${BASE_URL}/login`, user);
      const userResponse:IUserResponse = response.data;
      //mapper
      const logegUser=FromUserResponse(userResponse)
      return { data: logegUser, success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  Register: async (
    email: string,
    password: string
  ): Promise<ApiResponse<IUser>> => {
    const user: IUserRequest = { Email: email, Password: password };
    try {
      const response = await axios.post(`${BASE_URL}/register`, user);
      const userResponse:IUserResponse = response.data;
      //mapper
      const logegUser=FromUserResponse(userResponse)
      return { data: logegUser, success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetUserById: async (id): Promise<ApiResponse<IUser>> => {
    try {
      const response = await axios.get<IUserResponse>(`${BASE_URL}/${id}`);
      const userResponse:IUserResponse = response.data;
      //mapper
      const user=FromUserResponse(userResponse)
      return { success: true, data: user };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateUser: async (id, user: IUserRequest): Promise<ApiResponse<IUser>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, user);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
