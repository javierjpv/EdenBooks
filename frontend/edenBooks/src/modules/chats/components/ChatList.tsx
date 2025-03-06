import { useEffect, useState } from "react"
import { chatService } from "../services/chatService"
import { IChatDto } from "../interfaces/IChatDto"
import { ChatItem } from "./ChatItem"
import { Box, CircularProgress, List, Typography } from "@mui/material"

export const ChatList = () => {
    const [loading, setloading] = useState<boolean>(true)
    const [chats, setchats] = useState<IChatDto[]>([])
    // const [error, seterror] = useState<boolean>(false)
const fetchChats=async():Promise<void>=>{
    setloading(true)
    // seterror(false)
    const response=await chatService.GetChats()
    setloading(false)
    console.log(response.data)
    if (response.success&&response.data) {
        setchats(response.data)      
    }else{
        // seterror(true)
    }
    
  
}

useEffect(() => {
  fetchChats()
}, [])

  return (
    <Box sx={{ width: "100%", maxWidth: 400, bgcolor: "background.paper", mx: "auto", mt: 2 }}>
      <Typography variant="h6" sx={{ textAlign: "center", mb: 2 }}>
        Chats
      </Typography>

      {loading ? (
        <Box sx={{ display: "flex", justifyContent: "center", mt: 2 }}>
          <CircularProgress />
        </Box>
      ) : chats.length > 0 ? (
        <List>
          {chats.map((chat) => (
            <ChatItem key={chat.ID} chat={chat} />
          ))}
        </List>
      ) : (
        <Typography variant="body2" sx={{ textAlign: "center", color: "text.secondary" }}>
          No hay chats disponibles
        </Typography>
      )}
    </Box>
  )
}
