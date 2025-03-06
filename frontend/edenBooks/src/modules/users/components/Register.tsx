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

export const Register = () => {
  //Creo un useState para saber si el usuario ha pulsado submit en este componente
  const [showError, setshowError] = useState(false);
  const [submitted, setsubmitted] = useState(false);
  //Uso mi customHook para los formularios
  const { email, password, handleSubmit, handleOnChange } = useForm({
    email: "",
    password: "",
  });

  //Extraigo startRegister (funcion la cual realiza el registro si los argumentos son correctos,
  //en caso contrario,hace que el errror se muestre en el formulario)
  const { startRegister, user } = useAuthStore();

  //Tras realizar el submit se intenta registar al usuario
  const accionTrasSubmit: () => Promise<void> = async () => {
    try {
      setsubmitted(true)
      await startRegister(email, password);
    } catch (error) {
      console.log("Ha ocurrido un error tras el submit de register");
    }
  };

  //Esta funcion es la que se llama tras realizar el submit
  const onHandleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    handleSubmit(e, accionTrasSubmit);
  };
  useEffect(() => {
    if (user?.Error.trim().length > 0&&submitted) {
      setshowError(true);
      setTimeout(() => {
        setshowError(false);
      }, 3000);
    }
  }, [user.Error,submitted]);

  return (
    <Container maxWidth="sm">
      <Paper elevation={3} sx={{ padding: 4, marginTop: 5 }}>
        <Typography variant="h4" align="center" gutterBottom>
          Register
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
              Crear Cuenta
            </Button>
          </Box>
        </form>
      </Paper>
    </Container>
  );
};
