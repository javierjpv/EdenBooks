import { addressApi } from "../api/addressApi";
import { IAddressRequest } from "../interfaces/IAddressRequest";
import { IAddressService } from "../interfaces/IAddressService";

export const AddressService: IAddressService = {
  CreateAddress: async (address: IAddressRequest) =>addressApi.CreateAddress(address),
  GetAddresss: async () => addressApi.GetAddresss(),
  GetAddressById: async (id: number) => addressApi.GetAddressById(id),
  UpdateAddress: async (id: number, address: IAddressRequest) =>addressApi.UpdateAddress(id, address),
  DeleteAddress: async (id: number) => addressApi.DeleteAddress(id),
};
