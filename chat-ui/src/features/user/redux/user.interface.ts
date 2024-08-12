type Status = 'online' | 'offline';

export interface UserMetadata {
  userId: string;
  email: string;
  displayName: string;
  friends: FriendHash;
}

export interface Friend {
  userId: string;
  email: string;
  displayName: string;
  isOnline: boolean;
}

export interface FriendHash {
  [key: string]: Friend;
}

export interface UserPresence {
  userId: string;
  status: Status;
}
