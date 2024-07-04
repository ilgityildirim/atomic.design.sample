<template>
  <div v-bind="$attrs">
    <h3 class="font-bold mb-2" v-if="!noTitle">
      {{ t('title') }}
    </h3>
    <div :class="[ optionsWrapperClass ? optionsWrapperClass : 'space-y-2' ]">
      <!--suppress JSUnresolvedVariable -->
      <BaseOption
          :placeholder="t('duration')"
          :options="state.durations.value"
          v-model="state.duration.value"
          :disabled="store.car.fuelType === ''"
          :class="optionClass || ''"
      />
      <!--suppress JSUnresolvedVariable -->
      <BaseOption
          :placeholder="t('distance')"
          :options="state.distances.value"
          v-model="state.distance.value"
          :disabled="store.premium.duration === ''"
          :class="optionClass || ''"
          ref="distance"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, WritableComputedRef } from 'vue';
import { Ref } from '@vue/reactivity';
import { useI18n } from 'vue-i18n';
import { data, Unlimited } from '@/config';
import {
  AdvancedPackage,
  BasePackage,
  BasicPackage,
  ComfortPackage,
  PackageLanguages,
  PackageName,
  PackagePricing, UltimatePackage,
} from '@/config/types';
import { Option } from '@/config/forms';
import { useCalculatorStore } from '@/store';
import BaseOption from '@/components/atoms/BaseOption.vue';

interface Props {
  noTitle?: boolean;
  optionsWrapperClass?: string;
  optionClass?: string;
  autoSelect?: boolean;
}

interface State {
  durations: Ref<Option[]>;
  distances: Ref<Option[]>;
  duration: WritableComputedRef<string>;
  distance: WritableComputedRef<string>;
}

interface MonthlyDistances {
  months: number;
  distances: string[];
}

const props = defineProps<Props>();
const { t, locale } = useI18n();
const store = useCalculatorStore();
const distanceRef = ref<HTMLInputElement|null>(null);
const filteredData = computed(() => {
  const filteredData: { [key: number]: string[] } = {};
  if (store.car.series === '' || store.car.fuelType === '') {
    return filteredData;
  }

  const isComfort = (obj: BasicPackage|ComfortPackage|AdvancedPackage|UltimatePackage) => {
    return 'monthly' in obj && 'upfront' in obj;
  };

  const getMonthsDistances = (obj: PackagePricing): MonthlyDistances[] => {
    const data: MonthlyDistances[] = [];
    if (!obj) {
      return data;
    }
    Object.keys(obj).forEach((key) => {
      if (obj[key] === 0) {
        return;
      }
      const split = key.split('/');
      if (split.length !== 2 || split[0] === '' || split[1] === '') {
        return;
      }

      let distance = split[1];
      if (distance === '0') {
        return;
      }
      else if(distance === 'Unlimited'){
        distance = t('unlimited');
      }

      const months = parseInt(split[0], 10);
      const index = data.findIndex(item => item.months === months);

      if (index === -1) {
        data.push({
          months,
          distances: [distance],
        });
      } else if (!data[index].distances.includes(distance)) {
        data[index].distances.push(distance);
      }
    });
    return data;
  };

  const concatMonthlyDistances = (monthlyDistances: MonthlyDistances[]) => {
    monthlyDistances.forEach((obj) => {
      if (typeof filteredData[obj.months] === 'undefined') {
        filteredData[obj.months] = obj.distances;
      } else {
        obj.distances.forEach((distance) => {
          if (!filteredData[obj.months].includes(distance)) {
            filteredData[obj.months].push(distance);
          }
        });
      }
    });
  };

  [PackageName.Basic, PackageName.Comfort, PackageName.Advanced, PackageName.Ultimate].forEach((key: PackageName) => {
    const items = (data[key] as BasePackage[])
        .filter(item => {
          return item.model === store.car.series ||
              item.fuelType[locale.value as PackageLanguages] === store.car.fuelType
              ;
        }) as ComfortPackage[]|BasicPackage[]|AdvancedPackage[]|UltimatePackage[]
    ;

    items.forEach((item) => {
      if (!isComfort(item)) {
        const monthlyDistances = getMonthsDistances((item as BasicPackage).price);
        concatMonthlyDistances(monthlyDistances);
        return;
      }

      const monthlyDistances = getMonthsDistances((item as ComfortPackage).monthly);
      concatMonthlyDistances(monthlyDistances);
    });
  });

  return filteredData;
});

const state: State = {
  durations: computed<Option[]>(() => {
    const durations:Option[] = [];
    Object.keys(filteredData.value)
        .sort((a, b) => parseInt(a) - parseInt(b))
        .forEach((key) => {
          durations.push({
            label: t('months', { months: key }),
            value: key,
          });
        })
    ;
    return durations;
  }),
  distances: ref([]),
  duration: computed({
    get: () => store.premium.duration,
    set: (value: string) => {
      store.setPremiumDuration(value, !props.autoSelect);
    },
  }),
  distance: computed({
    get: () => store.premium.distance,
    set: (value: string) => {
      if (value === t('unlimited')){
        value = Unlimited;
      }
      store.setPremiumDistance(value);
    },
  }),
};

watch(() => store.premium.duration, (newValue, oldValue) => {
  if (newValue === oldValue) {
    return;
  }

  if (!props.autoSelect) {
    state.distance.value = '';
  }
  state.distances.value = [
    ...filteredData.value[parseInt(newValue, 10)]
        .sort((a, b) => {
          // Convert strings to numbers if possible
          const numA = parseInt(a);
          const numB = parseInt(b);

          // Check if both values are numbers
          if (!isNaN(numA) && !isNaN(numB)) {
            return numA - numB; // Sort numbers in ascending order
          } else if (!isNaN(numA)) {
            return -1; // Numbers come before strings
          } else if (!isNaN(numB)) {
            return 1; // Strings come after numbers
          } else {
            // If both values are strings, do regular string comparison
            return a.localeCompare(b);
          }
        })
        .map((distance: string) => {
          let value = distance;
          let label = t('km', { distance: distance.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.') });
          if (distance === t('unlimited')){
            label = t('unlimited');
            value = Unlimited;
          }
          return {
            label,
            value,
          };
        }),
  ];

  if (props.autoSelect) {
    state.distance.value = state.distances.value[0].value;
  }

  if (distanceRef.value !== null) {
    distanceRef.value.focus();
  }
});
watch(() => store.car.fuelType, (newValue, oldValue) => {
  if (newValue === oldValue) {
    return;
  }
  state.duration.value = '';
  if (distanceRef.value !== null) {
    distanceRef.value.focus();
  }
});

if (store.premium.duration) {
  state.distances.value = [
    ...filteredData.value[parseInt(store.premium.duration, 10)]
        .sort((a, b) => parseInt(a) - parseInt(b))
        .map((distance: string) => {
          let value = distance;
          let label = t('km', { distance: distance.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.') });
          if (distance === t('unlimited')){
            label = t('unlimited');
            value = Unlimited;
          }
          return {
            label,
            value,
          };
        }),
  ];
}
</script>

<!--suppress SpellCheckingInspection -->
<i18n>
fr:
  title: 'Détails de la prime :'
  duration: Durée souhaitée
  distance: Kilométrage souhaité
  months: '{months} mois'
  km: '{distance} km'
  unlimited: Illimité
nl:
  title: 'Details premie:'
  duration: Gewenste looptijd
  distance: Gewenst aantal kilometers
  months: '{months} maanden'
  km: '{distance} km'
  unlimited: Onbeperkt
de:
  title: 'Details premie:'
  duration: Gewenste looptijd
  distance: Gewenst aantal kilometers
  months: '{months} maanden'
  km: '{distance} km'
  unlimited: Onbeperkt
</i18n>