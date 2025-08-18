import networkGenerator from '@/services/_config/network-generator';
import type {
  GetProjectResItem,
  LikeProjectBody,
  ProjectFormBody,
  ProjectTypeItem,
} from '@/services/projects/projects.types';
import { useMutation, useQuery } from '@tanstack/react-query';

const apiPrefix = '/projects';

export default function useProject() {
  const useListProject = () =>
    useQuery<GetProjectResItem>({
      queryKey: ['projects'],
      queryFn: () => networkGenerator(apiPrefix, 'get'),
    });

  const useDetailProject = (id: string | number) =>
    useQuery({
      queryKey: ['projects', id],
      queryFn: () =>
        networkGenerator<ProjectTypeItem>(`${apiPrefix}/${id}`, 'get'),
    });

  const createProject = useMutation({
    mutationFn: (body: ProjectFormBody) =>
      networkGenerator<GetProjectResItem, ProjectFormBody>(
        apiPrefix,
        'post',
        body
      ),
  });

  const updateProject = useMutation({
    mutationFn: ({ id, ...body }: ProjectFormBody & { id: number | string }) =>
      networkGenerator<GetProjectResItem, ProjectFormBody>(
        `${apiPrefix}/${id}`,
        'put',
        body
      ),
  });

  const deleteProject = useMutation({
    mutationFn: ({ id, ...body }: ProjectFormBody & { id: number | string }) =>
      networkGenerator<GetProjectResItem, ProjectFormBody>(
        `${apiPrefix}/${id}`,
        'delete',
        body
      ),
  });

  const likeProject = useMutation({
    mutationFn: ({ id, action }: LikeProjectBody & { id: number | string }) =>
      networkGenerator<GetProjectResItem, LikeProjectBody>(
        `${apiPrefix}/${id}/like`,
        'put',
        { action }
      ),
    onSettled() {
      // data, error, variables, context
    },
    onMutate({ action }) {
      return { action };
    },
    onError(_, { action }) {
      return { action: !action ? 1 : 0 };
    },
  });

  return {
    listProject: useListProject,
    detailProject: useDetailProject,
    createProject,
    updateProject,
    deleteProject,
    likeProject,
  };
}
