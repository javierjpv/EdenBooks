export interface IAddressDto {
    ID:number
    CreatedAt:string
    UpdatedAt: string
    DeletedAt: string|null
    City: string;
    Province: string;
    PostalCode: string;
    Country: string;
    Street: string;
    Number: number;
    Users:null
	Orders:null
  }