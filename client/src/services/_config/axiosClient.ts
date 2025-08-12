import axios, { AxiosError } from 'axios';
import delve from 'dlv';
import { siteConfig } from './site-config';

let is401Handled = false;

const axiosClient = axios.create({
  baseURL: siteConfig.apiBaseUrl,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
});

axiosClient.interceptors.request.use(function (config) {
  config.headers['Authorization'] = `Bearer ${localStorage.getItem('token')}`;
  return config;
});

axiosClient.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    if (error.response) {
      const status = delve(error, 'response.status');
      if (status === 401 && !is401Handled) {
        if (!window.localStorage.getItem('refresh_token')) {
          window.localStorage.clear();
          window.location.href = '/';
          return Promise.reject(error);
        }
        axios
          .post(
            `${import.meta.env.VITE_API_HOST}/v1/auth/refresh`,
            {
              refresh_token: window.localStorage.getItem('refresh_token'),
            },
            {
              headers: {
                uuid: window.navigator.userAgent,
              },
            }
          )
          .then((res) => {
            const token = delve(res, 'data.data.access_token');
            const refreshToken = delve(res, 'data.data.refresh_token');
            const role = delve(res, 'data.data.admin_role');
            window.localStorage.setItem('token', token);
            window.localStorage.setItem('refresh_token', refreshToken);
            window.localStorage.setItem('role', role);
            window.location.reload();
          })
          .catch((error: AxiosError) => {
            if (error.response) {
              const status = delve(error, 'response.status');
              const refreshTokenRes = delve(
                error,
                'response.data.errors.refresh_token'
              );
              if (refreshTokenRes === 'NOT_BLANK') {
                window.location.href = '/';
              }
              if (status === 401) {
                window.localStorage.clear();
                window.location.href = '/';
              }
            }
          });
        is401Handled = true;
      }
      if (status === 403) window.location.href = '/';
    }
    return Promise.reject(error);
  }
);

export default axiosClient;
