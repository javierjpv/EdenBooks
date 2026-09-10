import { NotificationList } from "../components/NotificationList";
import { Container } from "@mui/material";

export const NotificationPage = () => {
  return (
    <>
      <Container maxWidth={"md"} sx={{ marginTop: 22 }}>
        <NotificationList />
      </Container>
    </>
  );
};
