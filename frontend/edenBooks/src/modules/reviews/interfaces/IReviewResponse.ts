export interface IReviewResponse {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null | string;
  Rating: number;
  Comment: string;
  UserID: number;
  ProductID: number;
}
