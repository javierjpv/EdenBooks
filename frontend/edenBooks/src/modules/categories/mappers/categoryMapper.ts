import { ICategoryDto } from "../interfaces/ICategoryDto";
import { ICategory } from "../interfaces/ICategory";
export const FromCategoryDto = (categoryDto: ICategoryDto):ICategory => {
  let category: ICategory;
  category = {
    ID:categoryDto.ID,
    Name: categoryDto.Name,
    Description: categoryDto.Description,
  };
  return category
};
