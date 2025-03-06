import { IReview } from "../interfaces/IReview";
import { IReviewDto } from "../interfaces/IReviewDto";

export const fromDto = (reviewDto: IReviewDto): IReview => {
  const review: IReview = {
    ID: reviewDto.ID,
    comment: reviewDto.Comment,
    rating: reviewDto.Rating,
    productID: reviewDto.ProductID,
    userID: reviewDto.UserID,
  };
  return review;
};
