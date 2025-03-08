import { IAddress } from "../interfaces/IAddress";
import { IAddressResponse } from "../interfaces/IAdressResponse";

export const fromDto=(addressDto:IAddressResponse):IAddress=>{
return{
    ID:addressDto.ID,
    country:addressDto.Country,
    city:addressDto.City,
    province:addressDto.Province,
    postalCode:addressDto.PostalCode,
    street:addressDto.Street,
    number:addressDto.Number,
}
}