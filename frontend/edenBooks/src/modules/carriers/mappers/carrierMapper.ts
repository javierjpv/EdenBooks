import { ICarrier } from "../interfaces/ICarrier";
import { ICarrierDto } from "../interfaces/ICarrierDto";

export const CarrierFromDto=(carrierDto:ICarrierDto):ICarrier=>{
return{
    ID:carrierDto.ID,
    contact:carrierDto.Contact,
    name:carrierDto.Name
}
}