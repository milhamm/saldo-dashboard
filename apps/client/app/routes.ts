import { type RouteConfig, index, layout, route } from '@react-router/dev/routes'

export default [
  layout('features/auth/layout.tsx', [
    route('login', 'features/auth/login/page.tsx'),
    route('logout', 'features/auth/logout/page.tsx'),
  ]),
  layout('features/dashboard/layout.tsx', [index('features/dashboard/home/page.tsx')]),
] satisfies RouteConfig
