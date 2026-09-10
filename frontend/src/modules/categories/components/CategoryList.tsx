import { List, ListItem } from "@mui/material"
import { ICategory } from "../interfaces/ICategory"
import { ICategoryResponse } from "../interfaces/ICategoryResponse"
export const CategoryList = ({categories}:{categories:ICategory[]|ICategoryResponse[]}) => {
  return (
    <>
    <h2>CategoryList</h2>
    <List>
    {categories.map((category)=>(<ListItem key={category.ID}>{category.Name}</ListItem>))}
    </List>
    
    </>
  )
}
