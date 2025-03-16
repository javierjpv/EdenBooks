import {
  FormControl,
  Typography,
  Input,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
  Button,
  Alert,
  Stack,
  CircularProgress,
} from "@mui/material";

import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { IProduct } from "../interfaces/IProduct";
import { productService } from "../services/productService";
import { ICategory } from "../../categories/interfaces/ICategory";
import { useNavigate } from "react-router";
import { ArrowBackIosNew } from "@mui/icons-material";
import { useAuthStore } from "../../users/hooks/useAuthStore";
import axios from "axios";
import { IProductRequest } from "../interfaces/IProductResquest";
import { CategoryService } from "../../categories/services/categoryService";

const checkProduct = (product: IProductRequest): boolean => {
  if (
    product.Name.trim() === "" ||
    product.Description.trim() === "" ||
    product.Price <= 0 ||
    product.CategoryID <= 0 ||
    product.UserID <= 0
  ) {
    return false;
  }
  return true;
};

export const ProductForm = ({ id }: { id?: number }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const [loading, setloading] = useState(false); //Este loading es del submit
  //crear otro loading para rescuperar el producto en modo edicion
  const [error, seterror] = useState(false);
  const [success, setsuccess] = useState(false);
  const [categories, setcategories] = useState<ICategory[]>([]);
  const [product, setproduct] = useState<IProductRequest>({
    Name: "",
    Description: "",
    Price: 0,
    CategoryID: 0,
    UserID: Number(user?.ID) || 0,
    ImageURL: "",
  });

  const handleGoBack = () => {
    navigate(-1); // Esto te lleva a la página anterior
  };
  const showError = () => {
    seterror(true);
    setTimeout(() => {
      seterror(false);
    }, 3000);
  };
  const showSuccess = () => {
    setsuccess(true);
    setTimeout(() => {
      setsuccess(false);
    }, 6000);
  };
  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setproduct((prevProduct) => ({
      ...prevProduct,
      [name]: name === "Price" ? Number(value) || 0 : value,
    }));
  };

  const handleSelectChange = (e: SelectChangeEvent<number>) => {
    setproduct((prevProduct) => ({
      ...prevProduct,
      CategoryID: Number(e.target.value) || 0,
    }));
  };
  const handleResetProduct = () => {
    setproduct({
      ...product,
      Name: "",
      Description: "",
      Price: 0,
      CategoryID: 0,
      UserID: Number(user?.ID) || 0,
      ImageURL: "",
    });
  };
  
  const handleImageUpload = async (file: File) => {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("upload_preset", "javier");

    try {
      const response = await axios.post(
        import.meta.env.VITE_CLOUDINARY_URL, //url cloudinary
        formData
      );
      return response.data.secure_url;
    } catch (error) {
      console.error("Error al subir la imagen a Cloudinary", error);
      throw error;
    }
  };

  const handleProductSubmit = async (
    e: FormEvent<HTMLFormElement>
  ): Promise<void> => {
    e.preventDefault();
    if (!checkProduct(product)) {
      showError();
      console.log("Hay un campo que nos e ha rellenado correctamente");
      return;
    }
    let productToSubmit = { ...product };

    const fileInput = document.getElementById("ImageURL") as HTMLInputElement;
    const file = fileInput?.files?.[0];

    setloading(true);
    if (file) {
      try {
        const imageUrl = await handleImageUpload(file);
        productToSubmit = { ...productToSubmit, ImageURL: imageUrl };
      } catch (error) {
        console.error("Error uploading image:", error);
        setloading(false);
        showError();
        return;
      }
    }
    try {
      console.log("CONSOLE", productToSubmit);
      if (id === undefined) {
        await productService.CreateProduct(productToSubmit);
      } else {
        await productService.UpdateProduct(id, productToSubmit);
      }
      setTimeout(() => {
        handleResetProduct();
        showSuccess();
        setTimeout(() => {
          navigate("/products");
        }, 400);
      }, 700);
      seterror(false);
    } catch (error) {
      showError();
      console.log("Error de la api al crear un producto", error);
    } finally {
      setloading(false);
    }
  };

  const fetchCategories = async () => {
    const response = await CategoryService.GetCategories();
    // setloadingCategories(false);
    if (response.success && response.data) {
      setcategories(response.data);
      return;
    }
    console.log(response.error);
  };


  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchProduct = async (id: number): Promise<void> => {
    try {
      const fetchedProduct: IProduct = await productService.GetProductById(id);
      setproduct({
        CategoryID: fetchedProduct.CategoryID,
        Description: fetchedProduct.Description,
        Name: fetchedProduct.Name,
        ImageURL: fetchedProduct.ImageURL,
        Price: fetchedProduct.Price,
        UserID: Number(user?.ID) || 0,
      });
    } catch (error) {
      console.log("Error al obtener el producto para ser editado", error);
    }
  };
  useEffect(() => {
    if (id) {
      fetchProduct(id);
    }
  }, [id]);
  // // //este useEffect solo sera de pruebas
  useEffect(() => {
    console.log("estado actual de productRequest", product);
  }, [product]);

  return (
    <>
      <>
        {id === undefined ? (
          <Typography align="center" marginBottom={7} variant="h4" gutterBottom>
            Crear Nuevo Producto
          </Typography>
        ) : (
          <Typography align="center" marginBottom={7} variant="h4" gutterBottom>
            Editar Producto
          </Typography>
        )}
        <Button
          onClick={handleGoBack}
          startIcon={<ArrowBackIosNew />}
          sx={{ marginBottom: 3 }}
        >
          Volver
        </Button>
        <form onSubmit={handleProductSubmit}>
          <Stack spacing={3}>
            <FormControl fullWidth>
              <InputLabel htmlFor="Name">Name</InputLabel>
              <Input
                id="Name"
                name="Name"
                value={product.Name}
                onChange={handleInputChange}
              />
            </FormControl>

            <FormControl fullWidth>
              <InputLabel htmlFor="Description">Description</InputLabel>
              <Input
                id="Description"
                name="Description"
                value={product.Description}
                onChange={handleInputChange}
              />
            </FormControl>

            <FormControl fullWidth>
              <InputLabel htmlFor="Price">Price</InputLabel>
              <Input
                id="Price"
                name="Price"
                type="number"
                inputProps={{ min: 0 }}
                slotProps={{
                  
                }}
                value={product.Price}
                onChange={handleInputChange}
              />
            </FormControl>

            <FormControl fullWidth>
             
              <Select
                id="CategoryID"
                value={product.CategoryID || ""}
                onChange={handleSelectChange}
                displayEmpty
              >
                <MenuItem value="" disabled>
                  Seleccione una categoría
                </MenuItem>
                {categories.map((category) => (
                  <MenuItem key={category.ID} value={category.ID}>
                    {category.Name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl fullWidth>
              <Input id="ImageURL" name="ImageURL" type="file" />
            </FormControl>

            {id === undefined ? (
              <Button
                disabled={loading}
                variant="contained"
                type="submit"
                fullWidth
                startIcon={
                  loading ? (
                    <CircularProgress size={20} color="inherit" />
                  ) : null
                }
              >
                Crear Producto
              </Button>
            ) : (
              <Button
                disabled={loading}
                variant="contained"
                type="submit"
                fullWidth
                startIcon={
                  loading ? (
                    <CircularProgress size={20} color="inherit" />
                  ) : null
                }
              >
                Editar Producto
              </Button>
            )}

            {error && id === undefined && (
              <Alert severity="error">Error al crear el producto</Alert>
            )}
            {success && id === undefined && (
              <Alert severity="success">Producto creado correctamente</Alert>
            )}
            {error && id !== undefined && (
              <Alert severity="error">Error al editar el producto</Alert>
            )}
            {success && id !== undefined && (
              <Alert severity="success">Producto editado correctamente</Alert>
            )}
          </Stack>
        </form>
      </>
    </>
  );
};
