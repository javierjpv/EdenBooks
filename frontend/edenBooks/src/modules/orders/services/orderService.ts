import { IOrder } from "../interfaces/IOrder";
import { IOrderService } from "../interfaces/IOrderService";
import { orderApi } from "../api/orderApi";

export const orderService: IOrderService = {
  CreateOrder: async (order: IOrder) =>orderApi.CreateOrder(order),
  GetOrders: async () => orderApi.GetOrders(),
  GetOrderById: async (id: number) => orderApi.GetOrderById(id),
  UpdateOrder: async (id: number, order: IOrder) =>orderApi.UpdateOrder(id, order),
  DeleteOrder: async (id: number) => orderApi.DeleteOrder(id),
};
