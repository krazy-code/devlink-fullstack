import { useMutation, useQuery } from '@tanstack/react-query';
import type {
  DeveloperFormBody,
  DeveloperTypeItem,
  GetDeveloperResItem,
} from '../../services/developers/developers.types';
import networkGenerator from '../../services/network-generator';

const apiPrefix = '/developers';

export default function useDeveloper() {
  const useListDeveloper = () =>
    useQuery<GetDeveloperResItem>({
      queryKey: ['developers'],
      queryFn: () => networkGenerator(apiPrefix, 'get'),
    });

  const useDetailDeveloper = (id: string | number) =>
    useQuery({
      queryKey: ['developers'],
      queryFn: () =>
        networkGenerator<DeveloperTypeItem>(`${apiPrefix}/${id}`, 'get'),
    });

  const createDeveloper = useMutation({
    mutationFn: (body: DeveloperFormBody) =>
      networkGenerator<GetDeveloperResItem, DeveloperFormBody>(
        apiPrefix,
        'post',
        body
      ),
  });

  const updateDeveloper = useMutation({
    mutationFn: ({
      id,
      ...body
    }: DeveloperFormBody & { id: number | string }) =>
      networkGenerator<GetDeveloperResItem, DeveloperFormBody>(
        `${apiPrefix}/${id}`,
        'put',
        body
      ),
  });

  return {
    listDeveloper: useListDeveloper,
    detailDeveloper: useDetailDeveloper,
    createDeveloper,
    updateDeveloper,
  };
}
