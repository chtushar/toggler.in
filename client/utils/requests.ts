import axios from 'axios';
import type { AxiosRequestConfig } from 'axios';
import { ENV } from '../constants/env';

const instance = axios.create();

export const requests = {
  get: <T>(url: string): Promise<T> =>
    instance.get<T>(`${ENV.API_ROOT}${url}`).then(({ data }) => data),
  post: <T, Q>(
    url: string,
    body?: T,
    config?: AxiosRequestConfig,
  ): Promise<Q> =>
    instance
      .post<Q>(`${ENV.API_ROOT}${url}`, body, config)
      .then(({ data }) => data),
  put: <T, Q>(url: string, body?: T): Promise<Q> =>
    instance.put<Q>(`${ENV.API_ROOT}${url}`, body).then(({ data }) => data),
  delete: <T>(url: string): Promise<T> =>
    instance.delete<T>(`${ENV.API_ROOT}${url}`).then(({ data }) => data),
};
