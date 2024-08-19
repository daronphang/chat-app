type Status = 'online' | 'offline';

export interface UserMetadata {
  userId: string;
  email: string;
  displayName: string;
  recipients: RecipientHash;
}

export interface Recipient {
  userId: string;
  email: string;
  displayName: string;
  isFriend: boolean;
  friendName: string;
  isOnline: boolean;
  color: string;
}

export interface RecipientHash {
  [key: string]: Recipient;
}

export interface UserPresence {
  clientId: string;
  status: Status;
  targetId: string;
}
