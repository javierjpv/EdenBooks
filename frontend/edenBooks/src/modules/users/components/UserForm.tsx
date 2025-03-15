import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { useAuthStore } from "../hooks/useAuthStore";
import { userService } from "../services/userService";

import {
  Alert,
  Button,
  CircularProgress,
  FormControl,
  Input,
  InputLabel,
  Snackbar,
  Stack,
  Typography,
} from "@mui/material";
import { IUserRequest } from "../interfaces/IUserRequest";
import { IUserResponse } from "../interfaces/IUserResponse";
import axios from "axios";
import { useDispatch} from "react-redux";
import { onUpdateProfileImage } from "../../../Store/auth/authSlice";
// const checkUser = (user: IUserRequest): boolean => {
//   if (user.Email && user.Email.trim() === "") {
//     return false;
//   }
//   return true;
// };
export const UserForm = () => {
  const { user } = useAuthStore();
  
const dispatch = useDispatch();
  const [profileUserRequest, setprofileUserRequest] =
    useState<IUserRequest | null>(null);
  const [loadingProfileUser, setloadingProfileUser] = useState<boolean>(true);
  const [loadingSubmitUser, setloadingSubmitUser] = useState<boolean>(false)
  const [openSnackbar, setOpenSnackbar] = useState<boolean>(false);
  const fetchUser = async () => {
    console.log("fetching user")
    if (!user.ID) {
      setloadingProfileUser(false);
      return;
    }
    const response = await userService.GetUserById(user.ID);
    if (response.success && response.data) {
      console.log("VALOR EN FETCHING USER",response.data)
      const data: IUserResponse = response.data;

      setprofileUserRequest(prev => ({
        ...prev,
        Name: data.Name,
        Tel: data.Tel,
        ImageURL: data.ImageURL,
      }));

      setloadingProfileUser(false);
      console.log(data);
    } else {
      console.log("Error al llamar a fetchUser");
    }
  };

  useEffect(() => {
    fetchUser();
  }, []);

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setprofileUserRequest((prevUserRequest) => ({
      ...prevUserRequest,
      [name]: value,
    }));
  };

  const handleUserSubmit = async (
    e: FormEvent<HTMLFormElement>
  ): Promise<void> => {
    if (!user.ID) {
      console.log("user.ID vacio");
      return;
    }
    if (!profileUserRequest) {
      console.log("profileUserRequest vacio");
      return;
    }
    e.preventDefault();
    // if (!checkUser(profileUserRequest)) {
    //   console.log("Hay un campo que nos e ha rellenado correctamente");
    //   return;
    // }
    let userToSubmit = { ...profileUserRequest };

    const fileInput = document.getElementById("ImageURL") as HTMLInputElement;
    const file = fileInput?.files?.[0];
    if (file) {
      try {
        setloadingSubmitUser(true);
        const imageUrl = await handleImageUpload(file);
        userToSubmit = { ...userToSubmit, ImageURL: imageUrl };
      } catch (error) {
        console.error("Error uploading image:", error);
        setloadingSubmitUser(false);

        return;
      }
    }
    if (!userToSubmit.ImageURL) {
      return
    }
    console.log("handleUserSubmit, user Request a enviar", userToSubmit);
    setloadingSubmitUser(true)
    const response=await userService.UpdateUser(user.ID, userToSubmit);
    if (!response.success) {
      console.log("Ha habido un error actualizando el user")
    }else{
      setOpenSnackbar(true);
      dispatch(onUpdateProfileImage(userToSubmit.ImageURL));
    }
    setloadingSubmitUser(false)
    
  };


  const handleImageUpload = async (file: File) => {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("upload_preset", "javier");

    try {
      const response = await axios.post(
        import.meta.env.VITE_CLOUDINARY_URL, //url cloudinary
        formData
      );
      return response.data.secure_url;
    } catch (error) {
      console.error("Error al subir la imagen a Cloudinary", error);
      throw error;
    }
  };

  return (
    <>
      {loadingProfileUser ? (
        <Typography>Cargando...</Typography>
      ) : profileUserRequest ? (
        <>
          <form onSubmit={handleUserSubmit}>
            <Stack spacing={3}>
              <FormControl fullWidth>
                <InputLabel htmlFor="Name">Name</InputLabel>
                <Input
                  id="Name"
                  name="Name"
                   value={profileUserRequest?.Name || ""}
                  onChange={handleInputChange}
                />
              </FormControl>

              <FormControl fullWidth>
                <InputLabel htmlFor="Name">Telefono</InputLabel>
                <Input
                  id="Tel"
                  name="Tel"
                  value={profileUserRequest.Tel}
                  onChange={handleInputChange}
                />
              </FormControl>

              <FormControl fullWidth>
              <Input id="ImageURL" name="ImageURL" type="file" />
            </FormControl>

              <Button
                disabled={loadingSubmitUser}
                variant="contained"
                type="submit"
                fullWidth
                startIcon={
                  loadingSubmitUser ? <CircularProgress size={20} color="inherit" /> : null
                }
              >
                Editar User
              </Button>
            </Stack>
          </form>
                    <Snackbar
                      open={openSnackbar}
                      autoHideDuration={3000}
                      onClose={() => setOpenSnackbar(false)}
                      anchorOrigin={{ vertical: "top", horizontal: "right" }}
                    >
                      <Alert onClose={() => setOpenSnackbar(false)} severity="info">
                        ¡Perfil editado correctamente!
                      </Alert>
                    </Snackbar>
        </>
      ) : (
        <Typography>Ha ocurrido un error cargando el usuario</Typography>
      )}
    </>
  );
};
