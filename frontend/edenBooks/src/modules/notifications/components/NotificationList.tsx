import { useEffect, useState } from "react";
import { INotification } from "../interfaces/INotification";
import { notificationService } from "../services/notificationService";
import { NotificationItem } from "./NotificationItem";
import { Typography, Stack, Skeleton, Card, CardContent } from "@mui/material";
export const NotificationList = () => {
  const [notifications, setnotifications] = useState<INotification[]>([]);
  const [loading, setloading] = useState<boolean>(true);

  const fetchProducts = async (): Promise<void> => {
    const response = await notificationService.GetNotifications();
    setloading(false);
    if (response.success && response.data) {
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
      {loading ? (
        <Stack spacing={2}>
          {[...Array(9)].map((_, index) => (
            <Card key={index} variant="outlined">
              <CardContent>
                <Stack direction="row" spacing={2} alignItems="center">
                  <Skeleton variant="text" width="70%" height={20} />
                  <Skeleton variant="rounded" width={80} height={24} />
                </Stack>
              </CardContent>
            </Card>
          ))}
        </Stack>
      ) : (
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
      )}
    </>
  );
};
