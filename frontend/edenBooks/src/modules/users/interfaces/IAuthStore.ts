export interface IAuthStore {
  user: {
    userState: string;
    ID: null | number;
    Name: string;
    Email: string;
    Token: string;
    ImageURL: string;
    Error: string|undefined;
  };
}
