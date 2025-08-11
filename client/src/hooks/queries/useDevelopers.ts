import { useMutation, useQuery } from '@tanstack/react-query';
import type {
  DeveloperFormBody,
  DeveloperTypeItem,
  GetDeveloperResItem,
} from '../../services/developers/developers.types';
import networkGenerator from '../../services/network-generator';

const apiPrefix = '/developers';

export default function useDevelopers() {
  const useListDevelopers = () =>
    useQuery({
      queryKey: ['developers'],
      queryFn: () =>
        networkGenerator<null, GetDeveloperResItem>(apiPrefix, 'get'),
    });

  const useDetailDeveloper = (id: string | number) =>
    useQuery({
      queryKey: ['developers'],
      queryFn: () =>
        networkGenerator<null, DeveloperTypeItem>(`${apiPrefix}/${id}`, 'get'),
    });

  const createDevelopers = useMutation({
    mutationFn: (body: DeveloperFormBody) =>
      networkGenerator<DeveloperFormBody, GetDeveloperResItem>(
        apiPrefix,
        'post',
        body
      ),
  });

  const updateDevelopers = useMutation({
    mutationFn: ({
      id,
      ...body
    }: DeveloperFormBody & { id: number | string }) =>
      networkGenerator<DeveloperFormBody, GetDeveloperResItem>(
        `${apiPrefix}/${id}`,
        'put',
        body
      ),
  });

  return {
    listDeveloper: useListDevelopers,
    detailDeveloper: useDetailDeveloper,
    createDevelopers,
    updateDevelopers,
  };
}
