import axios, { AxiosError } from "axios";
import { ApiResponse } from "../../../interfaces/apiResponse";
import { IReviewApi } from "../interfaces/IReviewApi";
import { IReviewResponse } from "../interfaces/IReviewResponse";
import { IReview } from "../interfaces/IReview";
import { FromReviewResponse } from "../mappers/reviewMapper";

const BASE_URL = "http://localhost:6969/reviews";

export const reviewApi: IReviewApi = {
  GetReviews: async (): Promise<ApiResponse<IReview[]>> => {
    try {
      const response = await axios.get<IReviewResponse[]>(BASE_URL);
      const reviewsDto = response.data;
      const reviews: IReview[] = reviewsDto.map((reviewDto) =>
        FromReviewResponse(reviewDto)
      );

      return { success: true, data: reviews };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  GetReviewById: async (id): Promise<ApiResponse<IReview>> => {
    try {
      const response = await axios.get<IReviewResponse>(`${BASE_URL}/${id}`);
      const reviewDto = response.data;

      const review: IReview = FromReviewResponse(reviewDto);

      return { success: true, data: review };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  UpdateReview: async (id, review: IReview): Promise<ApiResponse<IReview>> => {
    try {
      await axios.put(`${BASE_URL}/${id}`, review);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
  DeleteReview: async (id): Promise<ApiResponse<IReview>> => {
    try {
      await axios.delete(`${BASE_URL}/${id}`);
      return { success: true };
    } catch (error) {
      const axiosError = error as AxiosError<{ error: string }>;
      return { success: false, error: axiosError.request?.response };
    }
  },
};
