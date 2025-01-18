import { createRequire } from 'node:module'
import typescript from '@rollup/plugin-typescript'
const require = createRequire(import.meta.url)
const packageJson = require('./package.json')

/** @type {import('rollup').RollupOptions} */
export default {
  input: 'src/index.ts',
  output: [
    {
      format: 'esm',
      file: packageJson.module,
      sourcemap: true,
    },
    {
      format: 'cjs',
      file: packageJson.main,
      name: packageJson.name,
      sourcemap: true,
    },
  ],
  treeshake: true,
  plugins: [
    typescript({
      outputToFilesystem: false,
      tsconfig: './tsconfig.json',
    }),
  ],
  external: [
    ...Object.keys(packageJson.dependencies),
    ...Object.keys(packageJson.dependencies),
    'react/jsx-runtime',
  ],
}
