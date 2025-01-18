import { resolve } from 'node:path'
import { reactRouter } from '@react-router/dev/vite'
import autoprefixer from 'autoprefixer'
import tailwindcss from 'tailwindcss'
import { defineConfig } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'

export default defineConfig({
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer],
    },
  },
  optimizeDeps: {
    // Do not optimize internal workspace dependencies.
    exclude: ['@saldo-dashboard/shared-ui'],
  },
  resolve: {
    alias: [
      {
        // Configure an alias for the package. So, we don't have to restart
        // the Vite server each time when the former is performed.
        find: '@saldo-dashboard/shared-ui',
        replacement: resolve(__dirname, '../../packages/shared-ui/src/index.ts'),
      },
    ],
  },
  plugins: [reactRouter(), tsconfigPaths()],
})
