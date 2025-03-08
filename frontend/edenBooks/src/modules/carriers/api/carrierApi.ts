import axios, { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { ICarrierApi } from "../interfaces/ICarrierApi";
import { ICarrier } from "../interfaces/ICarrier";
import { CarrierFromDto } from "../mappers/carrierMapper";
import { ICarrierDto } from "../interfaces/ICarrierDto";
import { ICarrierRequest } from "../interfaces/ICarrierRequest";

const BASE_URL = "http://localhost:6969/carriers";

export const carrierApi: ICarrierApi = {
  CreateCarrier: async ( carrier: ICarrierRequest): Promise<ApiResponse<ICarrier>> => {
    try {
      const response= await axios.post(BASE_URL, carrier);
      const carrierDto = response.data;
      const createdCarrier: ICarrier = CarrierFromDto(carrierDto);

      return { success: true,data:createdCarrier};
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetCarriers: async (): Promise<ApiResponse<ICarrier[]>> => {
    try {
      const response = await axios.get<ICarrierDto[]>(BASE_URL);
      const carriersDto = response.data;
      const carriers: ICarrier[] = carriersDto.map((carrierDto) =>
        CarrierFromDto(carrierDto)
      );

      return { success: true, data: carriers };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetCarrierById: async (id): Promise<ApiResponse<ICarrier>> => {
    try {
      const response = await axios.get<ICarrierDto>(`${BASE_URL}/${id}`);
      const carrierDto = response.data;

      const carrier: ICarrier = CarrierFromDto(carrierDto);

      return { success: true, data: carrier };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateCarrier: async (id, carrier: ICarrierRequest): Promise<ApiResponse<ICarrier>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, carrier);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteCarrier: async (id): Promise<ApiResponse<ICarrier>> => {
    try {
      await axios.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
