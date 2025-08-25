import { follow, getFollows, unfollow } from '@/services/follow';
import type { FollowRequest } from '@/services/follow/follow.types';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

export function useListFollows(follower_id: string) {
  return useQuery({
    queryKey: ['follows', follower_id],
    queryFn: () => getFollows(follower_id),
    enabled: !!follower_id,
  });
}

export function useFollowMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (body: FollowRequest) => follow(body),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['follows', variables.follower_id],
      });
    },
  });
}

export function useUnfollowMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (body: FollowRequest) => unfollow(body),
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({
        queryKey: ['follows', variables.follower_id],
      });
    },
  });
}
