import { Data } from '@/config/types';

declare module 'data.json' {
  import { Data } from '@/config/types';
  const value: Data;
  export default value;
}