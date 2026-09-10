import { reviewApi } from "../api/reviewApi";
import { IReview } from "../interfaces/IReview";
import { IReviewService } from "../interfaces/IReviewService";

export const ReviewService: IReviewService = {
  GetReviews: async () => reviewApi.GetReviews(),
  GetReviewById: async (id: number) => reviewApi.GetReviewById(id),
  UpdateReview: async (id: number, review: IReview) =>reviewApi.UpdateReview(id, review),
  DeleteReview: async (id: number) => reviewApi.DeleteReview(id),
};
