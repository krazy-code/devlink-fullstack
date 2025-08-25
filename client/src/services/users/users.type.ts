export type UserTypeItem = {
  id: string;
  name: string;
  email: string;
  bio?: string;
  avatar?: string;
  created_at: string;
};

export interface GetUserResItem {
  count: number;
  users: UserTypeItem[];
}

export type UserFormBody = {
  name: string;
  email: string;
  bio?: string;
  avatar?: File | null;
};
