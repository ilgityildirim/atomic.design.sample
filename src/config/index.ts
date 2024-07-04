import jsonStringData from './data.json?raw';
import { Data } from '@/config/types';

export const data: Data = JSON.parse(jsonStringData);

export const LUXEMBOURG = 'lu';
export const BELGIUM = 'be';
export const API_URL = import.meta.env.APP_API_URL;
export const COUNTRY = import.meta.env.APP_COUNTRY || BELGIUM;

export const ModelOrder = [
  '1',
  '2',
  '3',
  '4',
  '5',
  '6',
  '7',
  '8',
  'i',
  'M',
  'X',
  'Z',
];

export const IModel = 'i';
export const IModelSeriesOrder = [
  'i4',
  'i5',
  'i7',
  'i8',
  'iX1',
  'iX2',
  'iX3',
  'iX',
];

export const StorageNames = {
  language: 'language',
};

export const Unlimited = 'Unlimited';
