import { Outlet, redirect } from 'react-router'
import { getUserSession } from '~/services/auth.server'
import { AuthProvider } from '../auth/provider'
import type { Route } from './+types/layout'

export async function loader({ request }: Route.LoaderArgs) {
  const session = await getUserSession(request)
  const token = session.get('access_token')
  if (!token) throw redirect('/login')
  return token
}

export default function DashboardLayout({ loaderData }: Route.ComponentProps) {
  return (
    <AuthProvider token={loaderData}>
      <Outlet />
    </AuthProvider>
  )
}
