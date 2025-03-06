import { List, ListItem } from "@mui/material"
import { ICategory } from "../interfaces/ICategory"
import { ICategoryDto } from "../interfaces/ICategoryDto"
export const CategoryList = ({categories}:{categories:ICategory[]|ICategoryDto[]}) => {
  return (
    <>
    <h2>CategoryList</h2>
    <List>
    {categories.map((category)=>(<ListItem key={category.ID}>{category.Name}</ListItem>))}
    </List>
    
    </>
  )
}
