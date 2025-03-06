import { carrierApi } from "../api/carrierApi";
import { ICarrier } from "../interfaces/ICarrier";
import { ICarrierService } from "../interfaces/ICarrierService";

export const CarrierService: ICarrierService = {
  CreateCarrier: async (carrier: ICarrier) =>carrierApi.CreateCarrier(carrier),
  GetCarriers: async () => carrierApi.GetCarriers(),
  GetCarrierById: async (id: number) => carrierApi.GetCarrierById(id),
  UpdateCarrier: async (id: number, carrier: ICarrier) =>carrierApi.UpdateCarrier(id, carrier),
  DeleteCarrier: async (id: number) => carrierApi.DeleteCarrier(id),
};
