import { useEffect, useState } from "react";
import { INotification } from "../interfaces/INotification";
import { notificationService } from "../services/notificationService";
import { NotificationItem } from "./NotificationItem";
import { Typography, Stack } from "@mui/material";
export const NotificationList = () => {
  const [notifications, setnotifications] = useState<INotification[]>([]);

  const fetchProducts = async (): Promise<void> => {
    const response = await notificationService.GetNotifications();
    if (response.success&&response.data) {
        setnotifications(response.data);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  return (
    <>
   
        <Typography variant="h4" gutterBottom>
          Notifications
        </Typography>

        <Stack spacing={2}>
          {notifications.length > 0 ? (
            notifications.map((notification) => (
              <NotificationItem
                key={notification.ID}
                notification={notification}
              />
            ))
          ) : (
            <Typography variant="body1" color="textSecondary">
              No notifications available.
            </Typography>
          )}
        </Stack>
    
    </>
  );
};
