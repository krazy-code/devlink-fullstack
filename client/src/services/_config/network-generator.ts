import type { AxiosRequestConfig } from 'axios';
import axiosClient from './axiosClient';

export type NetworkTypes =
  | 'get'
  | 'post'
  | 'put'
  | 'delete'
  | 'putForm'
  | 'postForm';

export default async function networkGenerator<T = unknown, R = unknown>(
  path: string,
  type: NetworkTypes,
  reqConfig?: R & AxiosRequestConfig<R>
): Promise<T> {
  return axiosClient[type](path, reqConfig).then((res) => res.data.data);
}
