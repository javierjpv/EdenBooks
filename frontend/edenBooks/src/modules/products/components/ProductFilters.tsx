import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Grid2,
  Skeleton,
  ToggleButton,
  ToggleButtonGroup,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import { useEffect, useState } from "react";
import { useSearchParams } from "react-router";
import { CategoryService } from "../../categories/services/categoryService";
import { ICategory } from "../../categories/interfaces/ICategory";

interface ProductFiltersProps {
  setOpenDrawer: (open: boolean) => void;
}
export const ProductFilters: React.FC<ProductFiltersProps> = ({
  setOpenDrawer,
}) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [fetchedcategories, setfetchedcategories] = useState<
    null | ICategory[]
  >(null);
  const [loadingCategories, setloadingCategories] = useState(true);
  const [category, setCategory] = useState<string | null>(
    searchParams.get("category_id")
  );
  const theme = useTheme();
  const isMediumScreen = useMediaQuery(theme.breakpoints.up("md"));

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
      <Box padding={2}>
        <h3>Filtrar por categoría</h3>

        <Box display={"flex"} flexDirection={"row"}>

          {!isMediumScreen && (
            <Button
              variant="contained"
              color="inherit"
              onClick={() => setOpenDrawer(false)}
              fullWidth
              sx={{ mb: 2 }}
            >
              Cerrar
            </Button>
          )}

          <Button
            variant="contained"
            color="error"
            onClick={clearFilters}
            fullWidth
            sx={{ mb: 2 }}
          >
            Limpiar Filtros
          </Button>
        </Box>

        <ToggleButtonGroup
          value={category}
          exclusive
          onChange={handleCategoryChange}
          aria-label="Filtrar categoría"
          orientation="vertical"
          fullWidth
        >
          {loadingCategories ? (
            Array(8)
              .fill(null)
              .map((_, index) => (
                <Grid2 size={{ xs: 12 }} key={index}>
                  <Card>
                    <Skeleton variant="rectangular" width="100%" height={80} />
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
      </Box>
    </>
  );
};
