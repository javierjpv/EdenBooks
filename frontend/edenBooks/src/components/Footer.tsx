import { Box, Typography, Link, Container, Divider } from "@mui/material";

export const Footer = () => {
  return (
    <Box
      component="footer"
      marginTop={22}
      sx={{
        bgcolor: "background.paper",
        py: 3,
        px: 2,
        position: "relative",
        bottom: 0,
        width: "100%",
      }}
    >
      <Container maxWidth="lg">
        {/* Divider between content and bottom */}
        <Divider sx={{ mb: 2 }} />

        <Box
          display="flex"
          flexDirection="column"
          alignItems="center"
          justifyContent="center"
        >
          <Typography variant="body2" color="text.secondary" align="center">
            {"© "}
            {new Date().getFullYear()}{" "}
            <Link color="inherit">
              Edenbooks
            </Link>
            . Todos los derechos reservados.
          </Typography>

          <Box display="flex" gap={2} mt={2}>
            <Link href="/contact" color="inherit">
              Contacto
            </Link>
            <Link href="/about" color="inherit">
              Acerca de
            </Link>
            <Link href="/privacy" color="inherit">
              Política de Privacidad
            </Link>
          </Box>
        </Box>
      </Container>
    </Box>
  );
};

