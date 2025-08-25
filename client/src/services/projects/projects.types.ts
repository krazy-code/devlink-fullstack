export type ProjectTypeItem = {
  created_at: string;
  description: string;
  id: string;
  project_url: string;
  tech_stacks: string;
  title: string;
  user_id: string;
  user_name: string;
};

export interface GetProjectResItem {
  count: number;
  projects: ProjectTypeItem[];
}

export type ProjectFormBody = {
  id: string;
  user_id: string;
  user_name: string;
  title: string;
  description: string;
  tech_stacks: string;
  project_url: string;
  created_at: string;
};
export type LikeProjectBody = {
  action: 1 | 0;
};
