import { defineStore } from 'pinia';

export interface CalculatorState {
  isShowForm: boolean,
  isShowResults: boolean,
  car: {
    model: string,
    series: string,
    fuelType: string,
  },
  premium: {
    duration: string,
    distance: string,
  },
}

export const useCalculatorStore = defineStore('calculator', {
  state: (): CalculatorState => ({
    isShowForm: false,
    isShowResults: false,
    car: {
      model: '',
      series: '',
      fuelType: '',
    },
    premium: {
      duration: '',
      distance: '',
    },
  }),
  getters: {
    isSet: (state: CalculatorState) => {
      return state.car.series !== ''
        && state.car.fuelType !== ''
        && state.premium.duration !== ''
        && state.premium.distance !== ''
      ;
    },
  },
  actions: {
    setCarModel(model: string) {
      this.car.model = model;
      this.car.series = '';
      this.car.fuelType = '';
      this.premium.duration = '';
      this.premium.distance = '';
    },
    setCarSeries(series: string) {
      this.car.series = series;
      this.car.fuelType = '';
      this.premium.duration = '';
      this.premium.distance = '';
    },
    setCarFuelType(fuelType: string) {
      this.car.fuelType = fuelType;
      this.premium.duration = '';
      this.premium.distance = '';
    },
    setPremiumDuration(duration: string, resetDistance: boolean = true) {
      this.premium.duration = duration;
      if (resetDistance) {
        this.premium.distance = '';
      }
    },
    setPremiumDistance(distance: string) {
      this.premium.distance = distance;
    },
    toggleForm() {
      this.isShowForm = !this.isShowForm;
    },
    toggleResults() {
      if (!this.isSet) {
        return;
      }
      this.isShowResults = !this.isShowResults;
    },
    reset() {
      this.isShowForm = false;
      this.isShowResults = false;
      this.car.model = '';
      this.car.series = '';
      this.car.fuelType = '';
      this.premium.duration = '';
      this.premium.distance = '';
    },
  },
});