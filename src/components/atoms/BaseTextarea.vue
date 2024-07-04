<template>
  <label
      v-if="props?.label"
      :for="uid.toString()"
      :class="[
          'block text-sm font-medium',
          hasError ? 'text-red-700' : 'text-gray-700',
          props.labelClass
      ]"
 >
    {{ props.label || '' }}
  </label>
  <textarea
      v-bind="$attrs"
      :id="uid.toString()"
      :class="[
          'mt-1 block w-full shadow-sm sm:text-sm rounded-md',
          hasError ?
            'text-red-700 border-red-300 focus:ring-red-500 focus:border-red-500' :
            'border-gray-300 focus:ring-blue-500 focus:border-blue-500',
      ]"
      v-model="model"
      :placeholder="props.placeholder || props.label"
      :aria-describedby="hasError ? `${uid}-error` : undefined"
      :aria-invalid="hasError ? true : undefined"
  />
  <p
      v-if="hasError"
      class="text-red-700 text-xs italic mt-1"
      :id="`${uid}-error`"
      aria-live="assertive"
 >
    <span v-for="(msg, index) in props.errors" :key="index" class="block">
      {{ msg }}
    </span>
  </p>
</template>
<script setup lang="ts">
import { computed } from 'vue';
import UniqueID from '@/utils/uniqueID';
import { ModelValue } from '@/config/forms';

type ModelValues = string|number|string[]|undefined;
interface Props {
  label?: string;
  placeholder?: string;
  modelValue: ModelValue<ModelValues>;
  labelClass?: string;
  errors?: string[];
}

const emit = defineEmits(['update:modelValue']);
const props = defineProps<Props>();
const uid = UniqueID().getID();
const hasError = computed(() => props.errors && props.errors.length> 0);
const model = computed({
  get: () => props.modelValue as string,
  set: (value: string) => emit('update:modelValue', value),
});
</script>