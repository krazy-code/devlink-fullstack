import type { AxiosRequestConfig } from 'axios';
import axiosClient from './_config/axiosClient';

export type NetworkTypes = 'get' | 'post' | 'put' | 'delete';

export default async function networkGenerator<T = unknown, R = unknown>(
  path: string,
  type: NetworkTypes,
  reqConfig?: T & AxiosRequestConfig<T>
): Promise<R> {
  return axiosClient[type](path, reqConfig).then((res) => res.data.data);
}
