import { type RouteConfig, index, layout, route } from '@react-router/dev/routes'

export default [
  index('routes/home.tsx'),
  layout('./features/auth/layout.tsx', [
    route('login', './features/auth/login/page.tsx'),
    route('logout', './features/auth/logout/page.tsx'),
  ]),
] satisfies RouteConfig
