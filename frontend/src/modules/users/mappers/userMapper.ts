
import { IUser } from "../interfaces/IUser";
import { IUserResponse } from "../interfaces/IUserResponse";
export const FromUserResponse = (userResponse:IUserResponse) => {
  let user: IUser;
  user = {
    ID:userResponse.ID,
    Email:userResponse.Email,
    Token:userResponse.Token,
    Name:userResponse.Name,
    AddressID:userResponse.AddressID,
    Tel:userResponse.Tel,
    ImageURL:userResponse.ImageURL

  };
  return user
};
