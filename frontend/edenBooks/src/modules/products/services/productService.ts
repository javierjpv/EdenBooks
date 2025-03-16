import { productApi } from "../api/productApi";
import { IProduct } from "../interfaces/IProduct";
import { IProductRequest } from "../interfaces/IProductResquest";
import { IProductService } from "../interfaces/IProductService";
import { FromProductResponse } from "../mappers/productMapper";

export const productService: IProductService = {
    CreateProduct: async (product: IProductRequest) =>productApi.CreateProduct(product),
    GetProducts: async (params?: URLSearchParams) => productApi.GetProducts(params),
    GetProductById: async (id: number) => productApi.GetProductById(id),
    UpdateProduct: async (id: number, product: IProductRequest) =>productApi.UpdateProduct(id, product),
    DeleteProduct: async (id: number) => productApi.DeleteProduct(id),
    
  GetFavorites: async (): Promise<IProduct[]> => {
    try {
      const productsDto = await productApi.GetFavorites();
      const products: IProduct[] = productsDto.map((productDto) =>
        FromProductResponse(productDto)
      );
      return products;
    } catch (error) {
      console.error("Error fetching products:", error);
      throw new Error("Error fetching products");
    }
  },
  AddToFavorite: async (id: number): Promise<void> => {
    try {
      await productApi.AddToFavorite(id);
    } catch (error) {
      console.log("Error al añadir un favorito");
      throw new Error("Error al añadir un favorito");
    }
  },
  DeleteFromFavorite: async (id: number): Promise<void> => {
    try {
      await productApi.DeleteFromFavorites(id);
    } catch (error) {
      console.log("Error al borrar un favorito");
      throw new Error("Error al borrar un favorito");
    }
  },
};
