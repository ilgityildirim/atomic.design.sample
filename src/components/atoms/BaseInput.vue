<template>
  <template v-if="!isRadioOrCheckbox">
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
    <input
        :type="props?.type || 'text'"
        v-bind="$attrs"
        :id="uid"
        :class="[
          'mt-1 block w-full shadow-sm sm:text-sm rounded-md',
          hasError ?
            'text-red-700 border-red-300 focus:ring-red-500 focus:border-red-500' :
            'border-gray-300 focus:ring-blue-500 focus:border-blue-500',
      ]"
        :value="props.modelValue"
        @input="onInput($event)"
        :placeholder="props.placeholder || props.label"
        :aria-describedby="hasError ? `${uid}-error` : undefined"
        :aria-invalid="hasError ? true : undefined"
    />
  </template>
  <template v-else>
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
    <label class="mt-1 flex items-center space-x-2.5 cursor-pointer">
      <input
          :type="props?.type"
          :value="props.modelValue"
          v-bind="$attrs"
          :id="uid"
          :class="[
              'shadow-sm inn',
              props.type === InputType.radio ? 'rounded-full' : 'rounded-sm',
              hasError ?
                'text-red-700 border-red-300 focus:ring-red-500 focus:border-red-500' :
                'border-gray-300 focus:ring-blue-500 focus:border-blue-500',
          ]"
          @change="onInput($event)"
          :aria-describedby="hasError ? `${uid}-error` : undefined"
          :aria-invalid="hasError ? true : undefined"
      />
      <span>{{ placeholder }}</span>
    </label>
  </template>
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
import { ModelValue, RawModelValue } from '@/config/forms';

interface Props {
  label?: string;
  placeholder?: string;
  modelValue: ModelValue<RawModelValue>;
  labelClass?: string;
  errors?: string[];
  type?: string;
}

enum InputType {
  checkbox = 'checkbox',
  radio = 'radio',
}

const emit = defineEmits(['update:modelValue']);
const props = defineProps<Props>();
const uid = UniqueID().getID().toString();
const hasError = computed(() => props.errors && props.errors.length> 0);
const isRadioOrCheckbox = computed(() => {
  return [InputType.radio, InputType.checkbox].includes(props?.type as InputType || '')
});
const onInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit('update:modelValue', target.value);
};
</script>