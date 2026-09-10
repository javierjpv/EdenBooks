import { Box, Button, Container, Paper, Slide } from "@mui/material";
import { Login } from "../components/Login";
import { Register } from "../components/Register";
import { useAuthStore } from "../hooks/useAuthStore";
import { Navigate } from "react-router";
import { useState } from "react";

//Tengo una pagina de autorizacion la cual contiene 2 elementos (Login y Register) esta pagina y
//estos elemntos soplo se mostraran si el usuario no esta autenticado ,una vez haya pasado el login o el register con
//exito esto dejara de mostrarse y se le redirigira a la ruta /    (que es el Home)

export const Auth = () => {
  const { user } = useAuthStore();
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => setIsLogin(!isLogin);

  if (user.userState === "AUTHENTICATED") {
    return <Navigate to="/" />;
  }
  return (
    <>
      <Container sx={{ marginTop: 29 }}>
        {user.userState === "AUTHENTICATED" ? (
          <Navigate to="/" />
        ) : (
          <Paper
            elevation={3}
            sx={{
              padding: 4,
              maxWidth: 500,
              margin: "auto",
              overflow: "hidden",
              minHeight: 600,
            }}
          >
            <Box display="flex" justifyContent="center" mb={2}>
              <Button variant="outlined" onClick={toggleForm}>
                {isLogin ? "Ir a Registro" : "Ir a Login"}
              </Button>
            </Box>

            {/* Contenedor con posición relativa para evitar desalineación */}
            <Box sx={{ position: "relative" }}>
              <Slide
                direction="left"
                in={isLogin}
                mountOnEnter
                unmountOnExit
                timeout={400}
              >
                <Box sx={{ position: "absolute", width: "100%" }}>
                  <Login />
                </Box>
              </Slide>

              <Slide
                direction="right"
                in={!isLogin}
                mountOnEnter
                unmountOnExit
                timeout={400}
              >
                <Box sx={{ position: "absolute", width: "100%" }}>
                  <Register />
                </Box>
              </Slide>
            </Box>
          </Paper>
        )}
      </Container>
    </>
  );
};
