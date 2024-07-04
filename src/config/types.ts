export interface Model {
  id: number;
  label: string;
}

export enum PackageName {
  Basic = 'basic',
  Comfort = 'comfort',
  Advanced = 'advanced',
  Ultimate = 'ultimate',
}

export enum PackageLanguages {
  fr = 'fr',
  nl = 'nl',
  de = 'de',
}

export type PackagePricing = {
  [key: string]: number;
}

export interface BasePackage {
  typeCode: string;
  maker: string;
  model: string;
  series: {
    [key in keyof typeof PackageLanguages]: string;
  };
  fuelType: {
    [key in keyof typeof PackageLanguages]: string;
  },
}

export interface BasicPackage extends BasePackage {
  price: PackagePricing;
}

export interface ComfortPackage extends BasePackage {
  monthly: PackagePricing;
  upfront: PackagePricing;
}

export interface AdvancedPackage extends BasePackage {
  monthly: PackagePricing;
  upfront: PackagePricing;
}

export interface UltimatePackage extends BasePackage {
  monthly: PackagePricing;
  upfront: PackagePricing;
}

export interface Data {
  [PackageName.Basic]: BasicPackage[];
  [PackageName.Comfort]: ComfortPackage[];
  [PackageName.Advanced]: AdvancedPackage[];
  [PackageName.Ultimate]: UltimatePackage[];
}

export interface FuelType extends Model {}