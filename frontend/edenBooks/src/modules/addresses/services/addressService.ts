import { addressApi } from "../api/addressApi";
import { IAddress } from "../interfaces/IAddress";
import { IAddressService } from "../interfaces/IAddressService";

export const AddressService: IAddressService = {
  CreateAddress: async (address: IAddress) =>addressApi.CreateAddress(address),
  GetAddresss: async () => addressApi.GetAddresss(),
  GetAddressById: async (id: number) => addressApi.GetAddressById(id),
  UpdateAddress: async (id: number, address: IAddress) =>addressApi.UpdateAddress(id, address),
  DeleteAddress: async (id: number) => addressApi.DeleteAddress(id),
};
