import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    // tanstackRouter({
    //   target: 'react',
    //   autoCodeSplitting: true,
    //   generatedRouteTree: 'routeTree.gen',
    // }),
    react(),
  ],
  resolve: {
    alias: {
      '@': '/src',
      '@components': '/src/components',
      '@lib': '/src/lib',
    },
  },
});
