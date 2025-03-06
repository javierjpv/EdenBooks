import { productApi } from "../api/productApi";
import { IProduct } from "../interfaces/IProduct";
import { IProductService } from "../interfaces/IProductService";
import { FromDto } from "../mappers/productMapper";

export const productService: IProductService = {
  GetProducts: async (): Promise<IProduct[]> => {
    try {
      const productsDto = await productApi.GetProducts();
      const products: IProduct[] = productsDto.map((productDto) =>
        FromDto(productDto)
      );
      return products;
    } catch (error) {
      console.error("Error fetching products:", error);
      throw new Error("Error fetching products");
    }
  },
  GetProductById: async (id: number): Promise<IProduct> => {
    try {
      const productDto = await productApi.GetProductById(id);
      const product = FromDto(productDto);
      return product;
    } catch (error) {
      console.log("Error al obtener un producto");
      throw new Error("Error al obtener un producto");
    }
  },
  CreateProduct: async (product: IProduct): Promise<void> => {
    try {
      await productApi.CreateProduct(product);
    } catch (error: any) {
      if (error.response) {
        console.log("Error del servidor:", error.response.data);
      } else {
        console.log("Error de conexión:", error.message);
      }
      throw new Error("Error al crear un producto");
    }
  },
  UpdateProduct: async (id: number, product: IProduct): Promise<void> => {
    try {
      await productApi.UpdateProduct(id, product);
    } catch (error) {
      console.log("Error al actualizar un producto");
      throw new Error("Error al actualizar un producto");
    }
  },
  DeleteProduct: async (id: number): Promise<void> => {
    try {
      await productApi.DeleteProduct(id);
    } catch (error) {
      console.log("Error al borrar un producto");
      throw new Error("Error al borrar un producto");
    }
  },
};
