import { ApiResponse } from "../../../interfaces/apiResponse";
import { IReview } from "./IReview";
export interface IReviewApi{
    GetReviewById:(id:number)=>Promise<ApiResponse<IReview>>
    GetReviews:()=>Promise<ApiResponse<IReview[]>>
    UpdateReview:(id:number,review:IReview)=>Promise<ApiResponse<IReview>>
    DeleteReview:(id:number)=>Promise<ApiResponse<IReview>>
} 