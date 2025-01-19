import { destroyUserSession } from '~/services/auth.server'
import type { Route } from './+types/page'

export async function action({ request }: Route.ActionArgs) {
  return await destroyUserSession(request)
}

export async function loader({ request }: Route.LoaderArgs) {
  return await destroyUserSession(request)
}
