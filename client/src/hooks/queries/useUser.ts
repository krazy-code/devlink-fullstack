import type { GetUserResItem } from '@/services/users/users.type';
import { useQuery } from '@tanstack/react-query';
import networkGenerator from '../../services/network-generator';

const apiPrefix = '/users';

export default function useUser() {
  const useListUser = () =>
    useQuery<GetUserResItem>({
      queryKey: ['users'],
      queryFn: () => networkGenerator(apiPrefix, 'get'),
    });

  //   const useDetailUser = (id: string | number) =>
  //     useQuery({
  //       queryKey: ['Users'],
  //       queryFn: () =>
  //         networkGenerator<UserTypeItem>(`${apiPrefix}/${id}`, 'get'),
  //     });

  //   const createUser = useMutation({
  //     mutationFn: (body: UserFormBody) =>
  //       networkGenerator<GetUserResItem, UserFormBody>(
  //         apiPrefix,
  //         'post',
  //         body
  //       ),
  //   });

  //   const updateUser = useMutation({
  //     mutationFn: ({
  //       id,
  //       ...body
  //     }: UserFormBody & { id: number | string }) =>
  //       networkGenerator<GetUserResItem, UserFormBody>(
  //         `${apiPrefix}/${id}`,
  //         'put',
  //         body
  //       ),
  //   });

  return {
    listUser: useListUser,
    // detailUser: useDetailUser,
    // createUser,
    // updateUser,
  };
}
