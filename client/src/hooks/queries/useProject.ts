import networkGenerator from '@/services/network-generator';
import type {
  GetProjectResItem,
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

  return {
    listProject: useListProject,
    detailProject: useDetailProject,
    createProject,
    updateProject,
    deleteProject,
  };
}
