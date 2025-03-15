import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { IAuthStore } from "../../modules/users/interfaces/IAuthStore";

//Estos eventos se llevaran a cabo una vez la solicitud
//Hacia el backend es correcta
// Obtener datos del localStorage (si existen)
interface IUserStore {
  userState: 'NOT-AUTHENTICATED' | 'CHECKING' | 'AUTHENTICATED';
  ID: number | null;
  Name: string;
  Email: string;
  Token: string;
  ImageURL: string;
  Error: string;
}
  let initialState: IAuthStore = {
    user: JSON.parse(localStorage.getItem("user") || 
      JSON.stringify({
        userState: 'NOT-AUTHENTICATED',
        ID: null,
        Name: "",
        Email: "",
        Token: "",
        ImageURL: "", 
        Error: "",
      } as IUserStore)
    ),
  };

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    //Cuando se esta comprobando si la solicitud al backend es correcta
    onChecking: (state) => {
      state.user = {...state.user, userState: 'CHECKING'}
      state.user = {...state.user, Token:''}
      state.user = {...state.user, Email:''}
      state.user = {...state.user, ID:null}
      state.user = {...state.user, ImageURL:''}
      state.user = {...state.user, Error:''}
     
    },
    //Se llevara a cabo una vez la solicitud de login del backend es correcta
    onLogin: (state, action: PayloadAction<{ Token: string; Email: string; ID: number; ImageURL: string }>) => {
        state.user = {...state.user, userState: 'AUTHENTICATED'}
        state.user = {...state.user, Token: action.payload.Token}
        state.user = {...state.user, Email: action.payload.Email}
        state.user = {...state.user, ID: action.payload.ID}
        state.user = {...state.user,ImageURL : action.payload.ImageURL}
        state.user = {...state.user, Error:''}
        localStorage.setItem("user", JSON.stringify(state.user));
     
    },
    //se llevara a cabo una vez se ha realizado la solicitud al backend para hacer logout y ademas se ha borrado el Token Token del local storage del cliente
    onLogout: (state, action: PayloadAction<{ Error: string|undefined }>) => {
        state.user = {...state.user, userState: 'NOT-AUTHENTICATED'}
        state.user = {...state.user, Token:''}
        state.user = {...state.user, Email:''}
        state.user = {...state.user, ID:null}
        state.user = {...state.user, ImageURL:''}
        state.user = {...state.user, Error: action.payload.Error}
        localStorage.removeItem("user");
    
    },
    //Se llevara a cabo una vez la solicitud de registro de un nuevo usuario del backend es correcta
    onRegister: (state, action: PayloadAction<{ Token: string; Email: string;ID: number}>) => {
      state.user = {...state.user, userState: 'AUTHENTICATED'}
      state.user = {...state.user, Token: action.payload.Token}
      state.user = {...state.user, Email: action.payload.Email}
      state.user = {...state.user, ID: action.payload.ID}
      state.user = {...state.user, Error:''}
      localStorage.setItem("user", JSON.stringify(state.user));

  },
  onUpdateProfileImage: (state, action: PayloadAction<string>) => {
    state.user.ImageURL = action.payload; // Permite actualizar la imagen desde cualquier parte de la app
    localStorage.setItem("user", JSON.stringify(state.user));
  },
},

});

//se exponen las funciones para poder usarlas en mis customhooks (en la carpeta Hooks)
export const { onLogin,onChecking,onLogout,onRegister,onUpdateProfileImage} =
  authSlice.actions;
export default authSlice.reducer;
