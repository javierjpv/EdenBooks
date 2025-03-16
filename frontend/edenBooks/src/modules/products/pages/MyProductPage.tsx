import { Button, Container, Drawer, Grid2, useMediaQuery, useTheme } from "@mui/material";
import { ProductList } from "../components/ProductList";
import { ProductFilters } from "../components/ProductFilters";
import { useState } from "react";

export const MyProductPage = () => {
  const [openDrawer, setOpenDrawer] = useState<boolean>(false);
  const theme = useTheme();
  const isMediumScreen = useMediaQuery(theme.breakpoints.up("md"));
  const isMobileScreen = useMediaQuery(theme.breakpoints.down("sm"));
  return (
    <>
      <Container sx={{ marginTop:22 }}>
        <h1>My ProductPage</h1>

        {/* Botón para abrir los filtros solo en móviles */}
        {!isMediumScreen && (
          <Button
            variant="contained"
            color="info"
            onClick={() => setOpenDrawer(true)}
            sx={{ ml: 2 }}
          >
            Filtros
          </Button>
        )}

        <Grid2 container spacing={2}>
          {/* Filtros como columna en pantallas medianas y grandes */}
          {isMediumScreen ? (
            <Grid2 size={{ xs:12,md: 3 }}>
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

         
          <Grid2 size={{xs:12,md: 9 }}>
            <ProductList />
          </Grid2>
        </Grid2>
      </Container>
    </>
  );
};
