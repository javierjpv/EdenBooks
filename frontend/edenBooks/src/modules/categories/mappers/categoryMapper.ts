import { ICategoryResponse } from "../interfaces/ICategoryResponse";
import { ICategory } from "../interfaces/ICategory";
export const FromCategoryDto = (categoryDto: ICategoryResponse):ICategory => {
  let category: ICategory;
  category = {
    ID:categoryDto.ID,
    Name: categoryDto.Name,
    Description: categoryDto.Description,
  };
  return category
};
