import { ICarrier } from "../interfaces/ICarrier";
import { ICarrierResponse } from "../interfaces/ICarrierResponse";

export const CarrierFromDto=(carrierDto:ICarrierResponse):ICarrier=>{
return{
    ID:carrierDto.ID,
    contact:carrierDto.Contact,
    name:carrierDto.Name
}
}