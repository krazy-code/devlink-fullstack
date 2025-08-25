import type { GetUserResItem, UserTypeItem } from '@/services/users/users.type';
import { useMutation, useQuery } from '@tanstack/react-query';
import networkGenerator from '../../services/_config/network-generator';

const apiPrefix = '/users';

export default function useUser() {
  const useListUser = () =>
    useQuery<GetUserResItem>({
      queryKey: ['users'],
      queryFn: () => networkGenerator(apiPrefix, 'get'),
    });

  const useDetailUser = (id: string | number) =>
    useQuery({
      queryKey: ['Users'],
      queryFn: () =>
        networkGenerator<UserTypeItem>(`${apiPrefix}/${id}`, 'get'),
    });

  //   const createUser = useMutation({
  //     mutationFn: (body: UserFormBody) =>
  //       networkGenerator<GetUserResItem, UserFormBody>(
  //         apiPrefix,
  //         'post',
  //         body
  //       ),
  //   });

  const updateUser = useMutation({
    mutationFn: ({ id, body }: { id: string; body: FormData }) =>
      networkGenerator<GetUserResItem, FormData>(
        `${apiPrefix}/${id}`,
        'putForm',
        body
      ),
    onSuccess: (_, _v, ctx) => {
      (ctx as { onSuccess(): void })?.onSuccess?.();
    },
    onError: (err, _v, ctx) => {
      (ctx as { onError(err: unknown): void })?.onError?.(err);
      console.log('Context : ', ctx);
    },
  });

  return {
    listUser: useListUser,
    detailUser: useDetailUser,
    // createUser,
    updateUser,
  };
}
