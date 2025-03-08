import { IProductDto } from "../interfaces/IProductDto";
import { IProduct } from "../interfaces/IProduct";

export const FromDto = (productDto: IProductDto):IProduct=> {
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
    is_favorite: productDto.is_favorite
  }
  return product
};
