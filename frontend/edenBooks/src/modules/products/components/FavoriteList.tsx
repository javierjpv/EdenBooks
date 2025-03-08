
import { Card, CardActions, CardContent, Grid2, Skeleton } from "@mui/material";
import { IProduct } from "../interfaces/IProduct";
import { useEffect, useState } from "react";
import { productService } from "../services/productService";
import { FavoriteItem } from "./FavoriteItem";
export const FavoriteList = () => {
  const [products, setproducts] = useState<IProduct[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, seterror] = useState(false);
  const fetchProducts = async () => {
    try {
      const fetchedProducts = await productService.GetFavorites();
      seterror(false);
      setproducts(fetchedProducts);
    } catch (error) {
      seterror(true);
      console.error("Error al obtener los productos:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);
  return (
    <>
      <h1>FavoriteList</h1>
      <Grid2
        marginTop={11}
        container
        rowSpacing={6}
        columnSpacing={6}
        direction="row"
        sx={{
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        {loading && !error
          ? Array(8)
              .fill(null)
              .map((_, index) => (
                <Grid2 size={{ xs: 12, sm: 6, md: 4, lg: 3 }} key={index}>
                  <Card>
                    <Skeleton
                      variant="circular"
                      width={40}
                      height={40}
                      sx={{ margin: 2 }}
                    />
                    <Skeleton
                      variant="text"
                      width="80%"
                      sx={{ marginLeft: 2 }}
                    />
                    <Skeleton variant="rectangular" width="100%" height={270} />
                    <CardContent>
                      <Skeleton variant="text" width="60%" />
                      <Skeleton variant="text" width="40%" />
                    </CardContent>
                    <CardActions disableSpacing>
                      <Skeleton variant="circular" width={30} height={30} />
                      <Skeleton variant="circular" width={30} height={30} />
                      <Skeleton variant="circular" width={30} height={30} />
                      <Skeleton variant="circular" width={30} height={30} />
                    </CardActions>
                  </Card>
                </Grid2>
              ))
          : ""}
        {!loading && !error
          ? products.map((product) => (
              <Grid2  key={product.ID} size={{ xs: 12, sm: 6, md: 4, lg: 3 }}>
                <FavoriteItem  product={product} />
              </Grid2>
            ))
          : ""}
        {error && <h3>Ha ocurrido un error obteniendo los productos</h3>}
      </Grid2>
    </>
  );
};

