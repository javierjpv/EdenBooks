
import { userApi } from "../api/userApi";
import { IUserRequest } from "../interfaces/IUserRequest";
import { IUserService } from "../interfaces/IUserService";

export const userService:IUserService= {
  Login: async (email: string, password: string) => userApi.Login(email, password),
  Register: async (email: string, password: string) => userApi.Register(email, password),
  GetUserById: async (id: number) => userApi.GetUserById(id),
  UpdateUser: async (id: number, user: IUserRequest) =>userApi.UpdateUser(id, user),
  
  
};
