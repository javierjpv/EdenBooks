import { IAddress } from "../interfaces/IAddress";
import { IAddressDto } from "../interfaces/IAdressDto";

export const fromDto=(addressDto:IAddressDto):IAddress=>{
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