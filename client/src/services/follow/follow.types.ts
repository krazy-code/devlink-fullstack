export type FollowItem = {
  id: string;
  follower_id: string;
  followed_id: string;
  created_at: string;
};

export type FollowRequest = {
  follower_id: string;
  followed_id: string;
};

export type ListFollowRes = {
  follows: FollowItem[];
};
