import {
  Avatar,
  Box,
  Divider,
  ListItemAvatar,
  ListItemText,
  Typography,
} from "@mui/material";
import { IChatResponse } from "../interfaces/IChatResponse";
import { Link } from "react-router";
import { AccountCircle } from "@mui/icons-material";
import { useAuthStore } from "../../users/hooks/useAuthStore";

export const ChatItem = ({ chat }: { chat: IChatResponse }) => {
  const { user } = useAuthStore();

  return (
    <>
      <Link
        to={`/chats/${chat.ID}`}
        style={{ textDecoration: "none", color: "inherit" }}
      >
        <Box display={"flex"} alignItems={"center"}>
          <ListItemAvatar>
            <Avatar>
              <AccountCircle></AccountCircle>
            </Avatar>
          </ListItemAvatar>
          {chat.Users.filter((chatUser) => chatUser?.ID != user.ID).map(
            (chatUser) => (
              <Typography key={chatUser.ID} variant="body2">{chatUser.Email}</Typography>
            )
          )}
        </Box>

        <ListItemText
          primary={`Chat #${chat.ID}`}
          secondary={new Date(chat.UpdatedAt).toLocaleString()}
        />
      </Link>

      <Divider />
      {chat.Messages?.map((message) => (
        <Typography key={message.ID} variant="body2">{message.Content}</Typography>
      ))}
    </>
  );
};
