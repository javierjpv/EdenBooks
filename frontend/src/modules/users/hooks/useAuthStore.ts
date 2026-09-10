import { useDispatch, useSelector } from "react-redux";
import {
  onChecking,
  onLogin,
  onLogout,
  onRegister,
} from "../../../Store/auth/authSlice";
import { userService } from "../services/userService";
import { RootState } from "../../../Store/store";
//Esto es un customHook para el estado global de la autenticacion
export const useAuthStore = () => {
  const { user } = useSelector((state: RootState) => state.auth);

  const dispatch = useDispatch();

  //Cuando se quiera hacer un login de un usuario se llama a esta funcion pasandole un Email y un Password
  const startLogin = async (Email: string, Password: string):Promise<void> => {
    //El estado del usuario se inicia en "CHECKING"
    console.log("CHECKING");
    dispatch(onChecking());
    //Se  realiza la solicitud al backend
    const response = await userService.Login(Email, Password);
    let data;
    console.log(response);
    //Si la respuesta del backend contiene un token  se guardara el token en el localStorage y  el estado del usuario pasara a "AUTHENTICATED"
    //a traves de redux
    if (response.success) {
      data = response.data;
      if (!data) {
        return
      }
      if (data!!.Token&&data!!.ID) {
        const Token = data.Token;
        const Email = data.Email;
        const ID=data.ID;
        const ImageURL=data.ImageURL??"";
        // localStorage.setItem("token", Token);
        console.log("Token received=>LOGIN: ", Token);
        dispatch(onLogin({ Token, Email,ID,ImageURL }));
      } else {
        console.log("falta token, id o imageURL en startLogin",response.data)
        //En caso contarrio se guardara el Error gracias a redux para poder ser usado y mostrarlo donde corresponda,ademas ,el estado del usuario
        //sera "NOT-AUTHENTICATED"
        const Error = response!!.error;
        dispatch(onLogout({ Error }));
        console.log("No token received:", response!!.error);
        // Manejar el caso en el que no se recibe un token del servidor
      }
    } else {
      const Error = response!!.error;
      dispatch(onLogout({ Error }));
    }
  };

  //Funcion para el registro de nuevos usuarios
  const startRegister = async (Email: string, Password: string):Promise<void> => {
    //El estado del usuario empezara en "CHECKING"
    console.log("CHECKING");
    dispatch(onChecking());
    //Se realizara la peticion al backend
    const response = await userService.Register(Email, Password);
    let data;
    if (response.success) {
      data = response.data;
      if (!data) {
        return
      }
      //Si la respuesta contiene un token este se guardara en el estado global del usuario ,ademas se guardara ese token en el localStorage y
      //el estado del usuario sera "AUTHENTICATED"
      if (data!!.Token&&data!!.ID) {
        const Token = data.Token;
        const Email = data.Email;
        const ID=data.ID;
    
        // const Email = data.Email;
        // localStorage.setItem("token", Token);
        console.log("Token received=>REGISTER: ", Token);
        dispatch(onRegister({ Token, Email,ID}));
      } else {
        //En caso contrario se guardara el Error recibido  y el estado "NOT-AUTHENTICATED" en el estado global del usuario (redux)
        const Error = response.error;
        dispatch(onLogout({ Error }));

        console.log("No token received:", response.error);

        // Manejar el caso en el que no se recibe un token del servidor
      }
    } else {
      console.log("LOGOUT");
      const Error = response.error;
      dispatch(onLogout({ Error }));

      // Manejar el Error de la solicitud al servidor
    }
  };

  //Funcion usada al hacer logout
  const startLogout = ():void => {
    console.log("LOGOUT");
    dispatch(onLogout({Error:""}));
  };

  //Se retorna un objeto con las funciones y las propiedades relacionadas con la autenticacion
  return {
    user,
    startLogin,
    startLogout,
    startRegister,
  };
};
