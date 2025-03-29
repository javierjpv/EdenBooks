import {
  Button,
  Box,
  Popover,
  List,
  ListItem,
  ListItemButton,
} from "@mui/material";
import { useSearchParams } from "react-router";
import { useState } from "react";
interface ProductSortOptionsProps {
  isMediumScreen: boolean;
}
export const ProductSortOptions = ({
  isMediumScreen,
}: ProductSortOptionsProps) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleSortChange = (order: "asc" | "desc", sortBy: "price") => {
    searchParams.set("order", order);
    searchParams.set("sort_by", sortBy);
    setSearchParams(searchParams);
    setAnchorEl(null); // Cierra el menú después de seleccionar
  };

  return (
    <Box display="flex" gap={2} >
      {!isMediumScreen && (
        <>
          <Button
            variant="contained"
            color="info"
            onClick={(e) => setAnchorEl(e.currentTarget)}
          >
            Ordenar
          </Button>
          <Popover
            open={open}
            anchorEl={anchorEl}
            onClose={() => setAnchorEl(null)}
            anchorOrigin={{ vertical: "bottom", horizontal: "left" }}
            transformOrigin={{ vertical: "top", horizontal: "left" }}
          >
            <List>
              <Button>Ordenar</Button>
              <ListItem disablePadding>
                <ListItemButton
                  onClick={() => handleSortChange("asc", "price")}
                >
                  Precio Ascendente
                </ListItemButton>
              </ListItem>
              <ListItem disablePadding>
                <ListItemButton
                  onClick={() => handleSortChange("desc", "price")}
                >
                  Precio Descendente
                </ListItemButton>
              </ListItem>
            </List>
          </Popover>
        </>
      )}
      {isMediumScreen && (
        <>
          <Button
            variant="contained"
            color="primary"
            onClick={() => handleSortChange("asc", "price")}
          >
            Precio Ascendente
          </Button>
          <Button
            variant="contained"
            color="secondary"
            onClick={() => handleSortChange("desc", "price")}
          >
            Precio Descendente
          </Button>
        </>
      )}
    </Box>
  );
};
