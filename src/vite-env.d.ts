/// <reference types="vite/client" />
declare const APP_VERSION: string;
declare const APP_API_URL: string;

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  // noinspection JSUnusedGlobalSymbols
  export default component
}
