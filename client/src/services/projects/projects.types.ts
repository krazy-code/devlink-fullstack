export type ProjectTypeItem = {
  created_at: string;
  description: string;
  id: string;
  project_url: string;
  tech_stacks: string;
  title: string;
  user_id: string;
};

export interface GetProjectResItem {
  count: number;
  projects: ProjectTypeItem[];
}

export type ProjectFormBody = {
  user_id: number;
  bio: string;
  github: string;
  location: string;
  website: string;
};
