import { IReview } from "../interfaces/IReview";
import {
  Card,
  CardContent,
  Typography,
  Paper,
  Box,
  Rating,
} from "@mui/material";

export const ReviewItem = ({ review }: { review: IReview | undefined }) => {
  return (
    <>
      {review !== undefined && (
        <>
          <Paper elevation={3} sx={{ padding: 2, borderRadius: 2 }}>
            <Card variant="outlined">
              <CardContent>
                <Box
                  display="flex"
                  justifyContent="space-between"
                  alignItems="center"
                >
                  <Typography variant="subtitle1" fontWeight="bold">
                    Usuario ID: {review.userID}
                  </Typography>
                  <Rating value={review.rating} readOnly />
                </Box>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ mt: 1 }}
                >
                  {review.comment}
                </Typography>
                <Typography variant="caption" color="text.disabled">
                  Producto ID: {review.productID}
                </Typography>
              </CardContent>
            </Card>
          </Paper>
        </>
      )}
    </>
  );
};
