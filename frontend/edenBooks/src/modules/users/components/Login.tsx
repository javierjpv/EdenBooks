import { FormEvent, useEffect, useState } from "react";
import { useForm } from "../../../hooks/useForm";
import { useAuthStore } from "../hooks/useAuthStore";
import {
  Box,
  Button,
  Container,
  Paper,
  TextField,
  Typography,
} from "@mui/material";

export const Login = () => {
  //Se usa este useState para saber si se ha hecho submit o no en este formulario
  const [showError, setshowError] = useState(false);
  const [submitted, setsubmitted] = useState(false);
  //Uso mi customHook para el manejo de formularios
  const { email, password, handleSubmit, handleOnChange } = useForm({
    email: "",
    password: "",
  });

  const { startLogin, user } = useAuthStore();

  const accionTrasSubmit = async (): Promise<void> => {
    try {
      //Se intenta realizar el login y si no se mostrara en este mismo componente el tipo de eroor que es (Se mostarra encima del boton de inicio sesion)
      //Fijate en que el error se mustra o no y eso lo sabremos por redux mediante user.error
      setsubmitted(true);
      await startLogin(email, password);
      console.log(user);
    } catch (error) {
      console.log(error);
    }
  };

  const onHandleSubmit = (e: FormEvent<HTMLFormElement>) => {
    handleSubmit(e, accionTrasSubmit);
  };

  useEffect(() => {
    if (user?.Error.length > 0 && submitted) {
      setshowError(true);
      setTimeout(() => {
        setshowError(false);
      }, 3000);
    }
  }, [user.Error, submitted]);

  return (
    <Container maxWidth="sm">
      <Paper elevation={3} sx={{ padding: 4, marginTop: 5 }}>
        <Typography variant="h4" align="center" gutterBottom>
          Login
        </Typography>
        <form onSubmit={onHandleSubmit}>
          <TextField
            fullWidth
            margin="normal"
            label="Email"
            variant="outlined"
            type="email"
            name="email"
            value={email}
            onChange={handleOnChange}
            placeholder="Introduce el email"
          />

          <TextField
            fullWidth
            margin="normal"
            label="Contraseña"
            variant="outlined"
            type="password"
            name="password"
            value={password}
            onChange={handleOnChange}
            placeholder="Introduce la contraseña"
            autoComplete="username"
          />

          {showError && (
            <Typography color="error" sx={{ mt: 1 }}>
              {user.Error}
            </Typography>
          )}

          <Box mt={2} display="flex" justifyContent="center">
            <Button type="submit" variant="contained" color="primary">
              Iniciar sesión
            </Button>
          </Box>
        </form>
      </Paper>
    </Container>
  );
};
