import { IProductResponse } from "../interfaces/IProductResponse";
import { IProduct } from "../interfaces/IProduct";

export const FromProductResponse = (productDto: IProductResponse):IProduct=> {
  let product: IProduct;
  product = {
    ID:productDto.ID,
    Name: productDto.Name,
    Description: productDto.Description,
    Price: productDto.Price,
    CategoryID: productDto.CategoryID,
    UserID: productDto.UserID,
    ImageURL: productDto.ImageURL,
    Sold:productDto.Sold,
    isFavorite: productDto.is_favorite
  }
  return product
};
