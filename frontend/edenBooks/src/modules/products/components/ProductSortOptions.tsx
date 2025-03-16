import { Button, Box } from "@mui/material";
import { useSearchParams } from "react-router";

export const ProductSortOptions: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const handleSortChange = (order: "asc" | "desc",sortBy: "price") => {
    searchParams.set("order", order);
    searchParams.set("sort_by", sortBy);
    setSearchParams(searchParams);
  };

  return (
    <Box display="flex" gap={2} sx={{ mb: 2 }}>
      <Button
        variant="contained"
        color="primary"
        onClick={() => handleSortChange("asc","price")}
      >
        Precio Ascendente
      </Button>
      <Button
        variant="contained"
        color="secondary"
        onClick={() => handleSortChange("desc","price")}
      >
        Precio Descendente
      </Button>
    </Box>
  );
};
