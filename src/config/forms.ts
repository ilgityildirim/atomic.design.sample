import { WritableComputedRef } from 'vue';
import { Ref } from '@vue/reactivity';

export interface Option {
  value: string;
  label: string;
  depth?: number;
}

export type RawModelValue = string | number | boolean | null | undefined;
export type ModelValue<T> = T | WritableComputedRef<T> | Ref<T>;
