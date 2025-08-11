export type DeveloperTypeItem = {
  id: number;
  name: string;
  github: string;
};

export interface GetDeveloperResItem {
  count: number;
  developers: DeveloperTypeItem[];
}

export type DeveloperFormBody = {
  user_id: number;
  bio: string;
  github: string;
  location: string;
  website: string;
};
