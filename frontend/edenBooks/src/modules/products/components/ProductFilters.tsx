import { Button, Card, CardActions, CardContent, Grid2, Skeleton, ToggleButton, ToggleButtonGroup } from "@mui/material";
import { useEffect, useState } from "react";
import { useSearchParams } from "react-router";
import { CategoryService } from "../../categories/services/categoryService";
import { ICategory } from "../../categories/interfaces/ICategory";

export const ProductFilters = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [fetchedcategories, setfetchedcategories] = useState<null | ICategory[]>(null);
  const [loadingCategories, setloadingCategories] = useState(true)
  const [category, setCategory] = useState<string | null>(searchParams.get("category_id"));

  const handleCategoryChange = (_: any, newCategory: string | null) => {
    if (newCategory) {
      searchParams.set("category_id", newCategory);
    } else {
      searchParams.delete("category_id");
    }
    setSearchParams(searchParams);
    setCategory(newCategory);
  };

  const clearFilters = () => {
    searchParams.delete("category_id");
    setSearchParams(searchParams);
    setCategory(null);
  };

  const fetchCategories = async () => {
    const response = await CategoryService.GetCategories();
    setloadingCategories(false);
    if (response.success && response.data) {
      setfetchedcategories(response.data);
      return;
    }
    console.log(response.error);
  };

  useEffect(() => {
    fetchCategories();
  }, [searchParams]);

  return (
    <>
      <h2>Filtrar por categoría</h2>
      <ToggleButtonGroup
        value={category}
        exclusive
        onChange={handleCategoryChange}
        aria-label="Filtrar categoría"
      >
        {/* <ToggleButton value="electronics">Electrónica</ToggleButton>
        <ToggleButton value="clothing">Ropa</ToggleButton>
        <ToggleButton value="home">Hogar</ToggleButton> */}
        {loadingCategories ? (
          Array(8)
          .fill(null)
          .map((_, index) => (
            <Grid2  size={{ xs: 12, sm: 6, md: 4, lg: 3 }} key={index}>
              <Card>
                <Skeleton variant="rectangular" width="100%" height={200} />
                <CardContent>
                  <Skeleton variant="text" width="60%" />
                </CardContent>
                <CardActions>
                  <Skeleton variant="circular" width={30} height={30} />
                  <Skeleton variant="circular" width={30} height={30} />
                </CardActions>
              </Card>
            </Grid2>
          ))
        ) : fetchedcategories ? (
          fetchedcategories.map((category) => (
            <ToggleButton key={category.ID} value={category.ID.toString()}>
              {category.Name}
            </ToggleButton>
          ))
        ) : (
          <p>No se pudieron cargar las categorías</p>
        )}
      </ToggleButtonGroup>

      <Button
        variant="contained"
        color="secondary"
        onClick={clearFilters}
        sx={{ ml: 2 }}
      >
        Limpiar Filtros
      </Button>
    </>
  );
};
