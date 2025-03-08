import { IReview } from "../interfaces/IReview";
import { IReviewResponse } from "../interfaces/IReviewResponse";

export const fromDto = (reviewDto: IReviewResponse): IReview => {
  const review: IReview = {
    ID: reviewDto.ID,
    comment: reviewDto.Comment,
    rating: reviewDto.Rating,
    productID: reviewDto.ProductID,
    userID: reviewDto.UserID,
  };
  return review;
};
