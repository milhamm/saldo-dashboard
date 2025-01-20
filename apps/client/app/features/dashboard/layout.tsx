import { Outlet } from 'react-router'
import type { Route } from './+types/layout'

export default function DashboardLayout(_: Route.ComponentProps) {
  return <Outlet />
}
