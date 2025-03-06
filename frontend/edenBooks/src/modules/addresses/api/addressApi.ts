import axios, { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IAddressApi } from "../interfaces/IAddressApi";
import { IAddress } from "../interfaces/IAddress";
import { fromDto } from "../mappers/addressMapper";
import { IAddressDto } from "../interfaces/IAdressDto";

const BASE_URL = "http://localhost:6969/addresses";

export const addressApi: IAddressApi = {
  CreateAddress: async ( address: IAddress): Promise<ApiResponse<IAddress>> => {
    try {
      const response= await axios.post(BASE_URL, address);
      const addressDto = response.data;
      const createdAddresss: IAddress = fromDto(addressDto);

      return { success: true,data:createdAddresss};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetAddresss: async (): Promise<ApiResponse<IAddress[]>> => {
    try {
      const response = await axios.get<IAddressDto[]>(BASE_URL);
      const addresssDto = response.data;
      const addresss: IAddress[] = addresssDto.map((addressDto) =>
        fromDto(addressDto)
      );

      return { success: true, data: addresss };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetAddressById: async (id): Promise<ApiResponse<IAddress>> => {
    try {
      const response = await axios.get<IAddressDto>(`${BASE_URL}/${id}`);
      const addressDto = response.data;

      const address: IAddress = fromDto(addressDto);

      return { success: true, data: address };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateAddress: async (id, address: IAddress): Promise<ApiResponse<IAddress>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, address);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteAddress: async (id): Promise<ApiResponse<IAddress>> => {
    try {
      await axios.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
