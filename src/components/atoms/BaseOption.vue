<template>
  <div
      :class="[
          '',
          props?.horizontal ? 'flex flex-row space-x-2 items-center' : 'flex flex-col',
      ]"
 >
    <label
        v-if="props?.label"
        :for="uid"
        :class="[
            'block text-sm font-medium',
            hasError ? 'text-red-700' : 'text-gray-700',
            props.labelClass
        ]"
   >
      {{ props.label || '' }}
    </label>
    <select
        v-bind="$attrs"
        :id="uid"
        :class="[
            'mt-1 block w-full shadow-sm sm:text-sm rounded-md px-3 py-4',
            hasError ?
              'text-red-700 border-red-300 focus:ring-red-500 focus:border-red-500' :
              'border-gray-300 focus:ring-blue-500 focus:border-blue-500',
            disabled ? 'opacity-50 cursor-not-allowed' : '',
        ]"
        @input="onSelect($event)"
        :aria-describedby="hasError ? `${uid}-error` : undefined"
        :aria-invalid="hasError ? true : undefined"
        :disabled="disabled ? true : undefined"
   >
      <option disabled v-if="props.label || props.placeholder" :selected="!modelValue">
        {{ props.label || props.placeholder }}
      </option>
      <option v-for="option in props.options" :value="option.value" :selected="modelValue === option.value">
        {{ option.label }}
      </option>
    </select>
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
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import UniqueID from '@/utils/uniqueID';
import { Option, ModelValue, RawModelValue } from '@/config/forms';

interface Props {
  label?: string;
  placeholder?: string;
  modelValue: ModelValue<RawModelValue>;
  options: Option[];
  labelClass?: string;
  errors?: string[];
  horizontal?: boolean;
  onChange?: (value: RawModelValue) => void;
  disabled?: boolean;
}

const emit = defineEmits(['update:modelValue']);
const props = defineProps<Props>();
const uid = UniqueID().getID().toString();
const hasError = computed(() => props.errors && props.errors.length> 0);
const onSelect = (event: Event) => {
  const value = (event.target as HTMLSelectElement).value;
  emit('update:modelValue', value);
  if (props?.onChange) {
    props.onChange(value);
  }
};
</script>