import {
  Box,
  Button,
  Container,
  Drawer,
  Grid2,
  Typography,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import { ProductList } from "../components/ProductList";
import { ProductFilters } from "../components/ProductFilters";
import { useState } from "react";
import { ProductSortOptions } from "../components/ProductSortOptions";

export const MyProductPage = () => {
  const [openDrawer, setOpenDrawer] = useState<boolean>(false);
  const theme = useTheme();
  const isMediumScreen = useMediaQuery(theme.breakpoints.up("md")); //El breakpoint de md lo veo mas util
  const isMobileScreen = useMediaQuery(theme.breakpoints.down("sm"));
  return (
    <>
      <Container sx={{ marginTop: 22 }}>
        <Typography textAlign={"center"} variant="h4">
          My ProductPage
        </Typography>
        <Box display={"flex"} gap={1.5}>
          {/* Botón para abrir los filtros solo en móviles */}
          {!isMediumScreen && (
            <Button
              variant="contained"
              color="info"
              onClick={() => setOpenDrawer(true)}
            >
              Filtrar
            </Button>
          )}
          {!isMediumScreen && (
            <ProductSortOptions isMediumScreen={isMediumScreen} />
          )}
        </Box>

        <Grid2 container spacing={2}>
          {/* Filtros como columna en pantallas medianas y grandes */}
          {isMediumScreen ? (
            <Grid2 size={{ xs: 12, md: 3 }}>
              <ProductFilters setOpenDrawer={setOpenDrawer} />
            </Grid2>
          ) : (
            /* Drawer en pantallas pequeñas (móviles) */
            <Drawer
              anchor="right"
              open={openDrawer}
              onClose={() => setOpenDrawer(false)}
              PaperProps={{
                sx: {
                  width: isMobileScreen ? "100%" : 320, // En xs ocupa toda la pantalla
                },
              }}
            >
              <ProductFilters setOpenDrawer={setOpenDrawer} />
            </Drawer>
          )}

          <Grid2 size={{ xs: 12, md: 9 }}>
            {isMediumScreen && (
              <ProductSortOptions isMediumScreen={isMediumScreen} />
            )}
            <ProductList />
          </Grid2>
        </Grid2>
      </Container>
    </>
  );
};
