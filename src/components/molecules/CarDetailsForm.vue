<template>
  <div v-bind="$attrs">
    <h3 class="font-bold mb-2">{{ t('title') }}</h3>
    <div class="space-y-2">
      <!--suppress JSUnresolvedVariable -->
      <BaseOption
          :options="state.models.value"
          v-model="state.model.value"
      />
      <!--suppress JSUnresolvedVariable -->
      <BaseOption
          :options="state.allSeries.value"
          v-model="state.series.value"
          :disabled="state.model.value === ''"
          ref="series"
      />
      <!--suppress JSUnresolvedVariable -->
      <BaseOption
          :options="state.fuelTypes.value"
          v-model="state.fuelType.value"
          :disabled="state.series.value === ''"
          ref="fuelType"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, WritableComputedRef } from 'vue';
import { Ref } from '@vue/reactivity';
import { useI18n } from 'vue-i18n';
import { data, IModel, IModelSeriesOrder, ModelOrder } from '@/config';
import { BasePackage, PackageLanguages, PackageName } from '@/config/types';
import { Option } from '@/config/forms';
import { useCalculatorStore } from '@/store';
import BaseOption from '@/components/atoms/BaseOption.vue';

interface State {
  models: Ref<Option[]>;
  allSeries: Ref<Option[]>;
  fuelTypes: Ref<Option[]>;
  model: WritableComputedRef<string>;
  series: WritableComputedRef<string>;
  fuelType: WritableComputedRef<string>;
}

const { t, locale } = useI18n();
const store = useCalculatorStore();
const fuelTypeRef = ref<HTMLInputElement|null>(null);
const seriesRef = ref<HTMLInputElement|null>(null);

const defaultFuelType: Option = {
  label: t('fuelType'),
  value: '',
};
const defaultSeries: Option = {
  label: t('series'),
  value: '',
};
const state: State = {
  models: computed<Option[]>(() => {
    const series: { label: string, value: string }[] = [
      {
        label: t('model'),
        value: '',
      },
    ];
    [PackageName.Basic, PackageName.Comfort, PackageName.Advanced, PackageName.Ultimate].forEach((key: PackageName) => {
      data[key].forEach((item) => {
        if (!series.some(obj => obj.value === item.model)) {
          series.push({ label: item.model, value: item.model });
        }
      });
    });

    series.sort((a, b) => {
      const aIndex = ModelOrder.indexOf(a.value);
      const bIndex = ModelOrder.indexOf(b.value);
      return aIndex - bIndex;
    });

    return series;
  }),
  allSeries: ref([defaultSeries]),
  fuelTypes: ref([defaultFuelType]),
  model: computed({
    get: () => store.car.model,
    set: (value: string) => {
      store.setCarModel(value);
    },
  }),
  series: computed({
    get: () => store.car.series,
    set: (value: string) => {
      store.setCarSeries(value);
    },
  }),
  fuelType: computed({
    get: () => store.car.fuelType,
    set: (value: string) => {
      store.setCarFuelType(value);
    },
  }),
};

watch(() => store.car.model, (newValue, oldValue) => {
  if (newValue === oldValue) {
    return;
  }

  state.series.value = '';
  state.allSeries.value = [defaultSeries];
  state.fuelType.value = '';
  state.fuelTypes.value = [defaultFuelType];

  [PackageName.Basic, PackageName.Comfort, PackageName.Advanced, PackageName.Ultimate].forEach((key: PackageName) => {
    (data[key] as BasePackage[]).filter(obj => obj.model === store.car.model).forEach((item) => {
      const series = item.series[locale.value as PackageLanguages] || '';
      if (!state.allSeries.value.some(obj => obj.value === series)) {
        state.allSeries.value.push({ label: series, value: series });
      }
    });
  });

  if (state.model.value === IModel) {
    state.allSeries.value.sort((a, b) => {
      const aVal = a.value.trim();
      const bVal = b.value.trim();
      const aIndex = IModelSeriesOrder.indexOf(aVal);
      const bIndex = IModelSeriesOrder.indexOf(bVal);
      return aIndex - bIndex;
    });
  }

  if (seriesRef.value !== null) {
    seriesRef.value.focus();
    if (state.series.value.length === 2) {
      state.series.value = state.allSeries.value[1].value;
    }
  }
});

watch(() => store.car.series, (newValue, oldValue) => {
  if (newValue === oldValue) {
    return;
  }

  state.fuelType.value = '';
  state.fuelTypes.value = [defaultFuelType];

  [PackageName.Basic, PackageName.Comfort, PackageName.Advanced, PackageName.Ultimate].forEach((key: PackageName) => {
    (data[key] as BasePackage[])
        .filter(item => item.model === state.model.value && item.series[locale.value as PackageLanguages] === newValue)
        .forEach((item) => {
          const fuelType = item.fuelType[locale.value as PackageLanguages] || '';
          if (!state.fuelTypes.value.some(obj => obj.value === fuelType)) {
            state.fuelTypes.value.push({
              label: fuelType,
              value: fuelType,
            });
          }
        });
    });

  if (fuelTypeRef.value !== null) {
    fuelTypeRef.value.focus();
  }
});
</script>

<!--suppress SpellCheckingInspection -->
<i18n>
fr:
  title: 'Détails de la voiture :'
  model: Modèle
  series: Série
  fuelType: Type de carburant
nl:
  title: 'Details wagen:'
  model: Model
  series: Reeks
  fuelType: Brandstoftype
de:
  title: 'Details wagen:'
  model: Model
  series: Reeks
  fuelType: Brandstoftype
</i18n>