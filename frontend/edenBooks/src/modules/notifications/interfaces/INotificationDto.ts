export interface INotificationDto {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    Content: string;
    Seen: boolean;
    UserID: number;
  }
