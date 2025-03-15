export interface IAuthStore {
  user: {
    userState: string;
    ID: null | number;
    Name: string;
    Email: string;
    Token: string;
    Error: string|undefined;
  };
}
