export interface IAuthStore {
  user: {
    userState: string;
    ID: null | number;
    Name: undefined|string;
    Email: string;
    Token: undefined|string;
    Error: string|undefined;
  };
}
