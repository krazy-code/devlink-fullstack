import axios from 'axios';
import type { FollowRequest, ListFollowRes } from './follow.types';

const API_PREFIX = '/api/v1/follows';

export const getFollows = async (follower_id: string) => {
  const { data } = await axios.get<ListFollowRes>(
    `${API_PREFIX}?follower_id=${follower_id}`
  );
  return data.follows;
};

export const follow = async (body: FollowRequest) => {
  return axios.post(API_PREFIX, body);
};

export const unfollow = async (body: FollowRequest) => {
  return axios.delete(API_PREFIX, { params: body });
};
