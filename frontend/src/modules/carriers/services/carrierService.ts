import { carrierApi } from "../api/carrierApi";
import { ICarrierRequest } from "../interfaces/ICarrierRequest";
import { ICarrierService } from "../interfaces/ICarrierService";

export const CarrierService: ICarrierService = {
  CreateCarrier: async (carrier: ICarrierRequest) =>carrierApi.CreateCarrier(carrier),
  GetCarriers: async () => carrierApi.GetCarriers(),
  GetCarrierById: async (id: number) => carrierApi.GetCarrierById(id),
  UpdateCarrier: async (id: number, carrier: ICarrierRequest) =>carrierApi.UpdateCarrier(id, carrier),
  DeleteCarrier: async (id: number) => carrierApi.DeleteCarrier(id),
};
