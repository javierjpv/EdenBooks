import { IOrderService } from "../interfaces/IOrderService";
import { orderApi } from "../api/orderApi";
import { IOrderRequest } from "../interfaces/IOrderRequest";

export const orderService: IOrderService = {
  CreateOrder: async (order: IOrderRequest) =>orderApi.CreateOrder(order),
  GetOrders: async () => orderApi.GetOrders(),
  GetOrderById: async (id: number) => orderApi.GetOrderById(id),
  UpdateOrder: async (id: number, order: IOrderRequest) =>orderApi.UpdateOrder(id, order),
  DeleteOrder: async (id: number) => orderApi.DeleteOrder(id),
};
