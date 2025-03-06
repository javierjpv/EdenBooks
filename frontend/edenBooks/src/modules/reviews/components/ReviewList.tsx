import { useEffect, useState } from "react";
import { IReview } from "../interfaces/IReview";
import { ReviewService } from "../services/reviewService";
import { ReviewItem } from "./ReviewItem";
import { Stack } from "@mui/material";




export const ReviewList = () => {

  const [reviews, setreviews] = useState<IReview[] | undefined>([]);
  const fetchReviews = async () => {
    const response = await ReviewService.GetReviews();
    if (response.success) {
      const fetchedReviews = response.data;
      setreviews(fetchedReviews);
    } else {
      console.log("Error al obtener las reviews");
    }
  };
  useEffect(() => {
    fetchReviews();
  }, []);

  return (
    <Stack marginTop={6} spacing={2} sx={{ width: "100%" }}>
      {reviews?.map((review) => (
        <ReviewItem key={review.ID} review={review} />
      ))}
    </Stack>
  );
};
