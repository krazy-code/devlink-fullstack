export type UserTypeItem = {
  id: string;
  name: string;
  email: string;
  created_at: string;
};

export interface GetUserResItem {
  count: number;
  users: UserTypeItem[];
}
