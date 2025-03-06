import { INotification } from "../interfaces/INotification";
import { Card, CardContent, Typography, Stack, Chip } from "@mui/material";

export const NotificationItem = ({ notification }: { notification: INotification }) => {
  return (
    <Card variant="outlined">
      <CardContent>
        <Stack direction="row" spacing={2} alignItems="center">
          <Typography variant="body1">{notification.Content}</Typography>
          <Chip label={`User: ${notification.UserID}`} color="primary" size="small" />
        </Stack>
      </CardContent>
    </Card>
  );
};
