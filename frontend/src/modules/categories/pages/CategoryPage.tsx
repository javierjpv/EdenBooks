import { useEffect, useState } from "react"
import { CategoryList } from "../components/CategoryList"
import { ICategory } from "../interfaces/ICategory"
import { categoryService } from "../services/categoryService"

export const CategoryPage = () => {
    const [categories, setcategories] = useState<ICategory[]>([])

    const fetchCategories=async():Promise<void>=>{
        const result:ICategory[]=await categoryService.GetCategories()
        setcategories(result)
    }
    useEffect(() => {
      fetchCategories()
    }, [])
    
  return (
    <>
    <h1>CategoryPage</h1>
    <CategoryList categories={categories}/>
    </>
  )
}
