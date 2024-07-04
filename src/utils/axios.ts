import axios, { AxiosRequestConfig } from 'axios';
import storage from '@/utils/storage';
import { API_URL, StorageNames } from '@/config';

export const defaultHeaders = {
  'Content-Type': 'application/json',
  'Access-Control-Allow-Origin': import.meta.env.APP_ALLOW_ORIGIN,
  'Accept': 'application/json',
}

export const headersConfig = (extras?: AxiosRequestConfig) => ({
  ...(Object.assign({}, extras, { headers: undefined }) || {}),
  headers: {
    ...defaultHeaders,
    'Accept-Language': `${storage.getItem(StorageNames.language) || ''}`,
    ...extras?.headers,
  },
});

const instance = axios.create({
  baseURL: API_URL,
  timeout: 3000,
  headers: defaultHeaders,
});

export default instance;