/// <reference types="vite/client" />

declare module 'asciinema-player' {
  export function create(
    url: string, 
    element: HTMLElement, 
    options?: {
      autoplay?: boolean;
      theme?: string;
      preload?: boolean;
      fit?: boolean;
      poster?: string;
    }
  ): void;
}
