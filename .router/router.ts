import { createRouter } from 'sv-router';
import About from '../src/routes/about.svelte';
import Index from '../src/routes/index.svelte';
import Layout from '../src/routes/layout.svelte';
import Settings from '../src/routes/settings.svelte';

export const routes = {
  '/about': About,
  '/': Index,
  'layout': Layout,
  '/settings': Settings
};
export type Routes = typeof routes;
export const { p, navigate, isActive, preload, route } = createRouter(routes);
