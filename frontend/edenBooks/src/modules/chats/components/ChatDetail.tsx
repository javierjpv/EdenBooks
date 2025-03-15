import {
  IconButton,
  Paper,
  Box,
  Container,
  Typography,
  ListItem,
  List,
  TextField,
  Button,
} from "@mui/material";
import { useNavigate, useParams } from "react-router";
import { chatService } from "../services/chatService";
import { useEffect, useRef, useState } from "react";
import { useAuthStore } from "../../users/hooks/useAuthStore";
import { ArrowBackIosNew, Send } from "@mui/icons-material";
import { IMessageResponse } from "../../messages/interfaces/IMessageResponse";
import { IMessageRequest } from "../../messages/interfaces/IMessageRequest";
import { IChat } from "../interfaces/IChat";

export const ChatDetail = () => {
  const { id } = useParams();
  const [loading, setloading] = useState<boolean>(true);
  const [chat, setchat] = useState<IChat | null>(null);
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const [newMessage, setNewMessage] = useState<string>("");
  const socket = useRef<WebSocket | null>(null);
  const loggedUserID: number = Number(user.ID);
  const [selectedSenderID, setselectedSenderID] = useState<number>(0);
  // const [error, seterror] = useState<boolean>(false)

  const fetchChat = async (): Promise<void> => {
    setloading(true);
    // seterror(false)
    const response = await chatService.GetChatById(Number(id));
    console.log(response.data);
    setloading(false);
    if (response.success && response.data) {
      setchat(response.data);
    } else {
      // seterror(true)
      console.log("error llamando a chatService.GetChatById");
    }
  };

  useEffect(() => {
    fetchChat();
  }, [id]);

  useEffect(() => {
    if (chat && chat.Users && user.ID !== null) {
      const validUsers = chat.Users.filter(
        (userFiltered) => userFiltered.ID !== undefined
      );
      console.log("USUARIOS VALIDADOS", validUsers);
      const foundUser = validUsers.find(
        (userFiltered) => userFiltered.ID != user.ID
      );
      console.log("Usuario encontrado", foundUser);
      if (foundUser && foundUser.ID) {
        const foundUserID = foundUser.ID;
        setselectedSenderID(foundUserID);
      } else {
        navigate(-1);
        //solucion temporal
        console.warn(
          "⚠ No se encontró un usuario válido para enviar el mensaje."
          //manejar este error, esto ocurre cuando hay un problema en la api
          //  al obtener los usuarios implicados en la conversacion
        );
      }
    }
  }, [chat]);

  useEffect(() => {
    socket.current = new WebSocket(
      `ws://localhost:8080/ws?token=${user.Token}`
    );
    socket.current.onopen = () => {
      console.log("WebSocket connected");
    };

    socket.current.onmessage = (event) => {
      console.log("Mensaje recibido:", event.data);
      const receivedMessage: IMessageResponse = JSON.parse(event.data);

      // Verificar que el senderID o el receiverID INCLUYA A AMBOS USUARIOS,
      //  EL LOGEADO Y EL OTRO MIEMBRO DEL CHAT
      //SI EL MENSAJE NO LO ENVIAN O NO LO RECIBEN NINGUN MIEMBRO DE ESTE CHAT
      //  SE IGNORARA

      console.log("OTRO MIEBRO DEL CHAT", selectedSenderID);

      if (
        receivedMessage.ReceiverID === user.ID ||
        receivedMessage.SenderID === user.ID ||
        receivedMessage.ReceiverID === selectedSenderID ||
        receivedMessage.SenderID == selectedSenderID
      ) {
        setchat((prevChat) => {
           if (!prevChat) return prevChat;

          return {
            ...prevChat,
            Messages: [...prevChat.Messages, receivedMessage],
          };
        });
      }
    };

    socket.current.onclose = () => {
      console.log("WebSocket disconnected. Reconnecting in 3 seconds...");
      setTimeout(() => {
        if (!socket.current || socket.current.readyState === WebSocket.CLOSED) {
          socket.current = new WebSocket(
            `ws://localhost:8080/ws?token=${user.Token}`
          );
        }
      }, 3000);
    };

    socket.current.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    return () => {
      if (socket.current) {
        socket.current.close();
      }
    };
  }, [selectedSenderID]);

  const sendMessage = () => {
    if (
      socket.current &&
      socket.current.readyState === WebSocket.OPEN &&
      newMessage.trim() !== ""
    ) {
      const messageData: IMessageRequest = {
        Content: newMessage,
        Seen: false,
        Status: "sent",
        ChatID: Number(id),
        SenderID: loggedUserID,
        ReceiverID: selectedSenderID,
      };
      console.log("MENSAJE ENVIADO", JSON.stringify(messageData));
      socket.current.send(JSON.stringify(messageData));
      setNewMessage("");
    }
  };
  const handleGoBack = () => {
    navigate(-1);
  };
  return (
    <Container sx={{ marginTop: 22 }}>
      {loading ? (
        <p>Cargando</p>
      ) : chat ? (
        <Box>
          <Typography variant="h3">{chat.ID}</Typography>
          {new Date(chat.UpdatedAt).toLocaleString()}
          {chat.Users !== null &&
            chat.Users.filter((chatUser) => chatUser?.ID != user.ID).map(
              (chatUser) => (
                <Box key={chatUser.ID}>
                   <Typography variant="body2">{chatUser.Name}</Typography>
                  <Typography variant="body2">{chatUser.Email}</Typography>
                </Box>
              )
            )}
          <Box />

          <Button
            onClick={handleGoBack}
            startIcon={<ArrowBackIosNew />}
            sx={{ marginBottom: 3 }}
          >
            Volver
          </Button>
          <Paper
            elevation={3}
            sx={{
              minHeight: { xs: 5, md: 400 },
              maxHeight: 570,
              flexGrow: 1,
              overflowY: "auto",
              overflowX: "hidden",
              padding: 2,
              borderRadius: 2,
              display: "flex",
              flexDirection: "column",
              width: "100%",
            }}
          >
            <List
              sx={{ flexGrow: 1, overflowY: "auto", width: "100%", gap: 1 }}
            >
              {chat.Messages  &&
                chat.Messages.map((message) => (
                  <ListItem
                    key={message.ID}
                    sx={{
                      display: "flex",
                      justifyContent:
                        message.SenderID === loggedUserID
                          ? "flex-end"
                          : "flex-start",
                      width: "100%",
                    }}
                  >
                    <Box
                      sx={{
                        maxWidth: "70%",
                        paddingY: 0.7,
                        paddingX: 4,
                        borderRadius: 4,
                        backgroundColor:
                          message.SenderID === loggedUserID
                            ? "#1976D2"
                            : "#E0E0E0",
                        color:
                          message.SenderID === loggedUserID ? "#FFF" : "#000",
                        wordBreak: "break-word",
                      }}
                    >
                      <Typography
                        variant="body1"
                        sx={{ whiteSpace: "pre-wrap" }}
                      >
                        {message.Content}
                      </Typography>
                    </Box>
                  </ListItem>
                ))}
            </List>
          </Paper>

          <Box sx={{ display: "flex", marginTop: 2, width: "100%" }}>
            <TextField
              label="Mensaje"
              variant="outlined"
              fullWidth
              value={newMessage}
              onChange={(e) => setNewMessage(e.target.value)}
              sx={{
                "& .MuiOutlinedInput-root": {
                  borderRadius: 4,
                  margin: 0,
                  padding: 0,
                },
                "& .MuiOutlinedInput-input": {
                  padding: "13px 10px",
                },
              }}
            />

            <IconButton onClick={sendMessage} sx={{ marginLeft: 1 }}>
              <Send />
            </IconButton>
          </Box>
        </Box>
      ) : (
        <p>No existe el chat</p>
      )}
    </Container>
  );
};
