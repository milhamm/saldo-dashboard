import { HydrationBoundary, QueryClient, dehydrate } from '@tanstack/react-query'
import { getAndValidateUserSession } from '~/services/auth.server'

import { ApiClient } from '~/services/api'
import type { Route } from './+types/page'

import { Movements } from './components/movements'
import { movementQueryOptions } from './hooks/use-movement'
import { MovementService } from './service'

export function meta() {
  return [
    { title: 'New React Router App' },
    { name: 'description', content: 'Welcome to React Router!' },
  ]
}

export async function loader({ request }: Route.LoaderArgs) {
  const token = await getAndValidateUserSession(request)

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 60 * 1000,
      },
    },
  })
  const apiClient = new ApiClient().setBearerToken(token)
  const movementService = new MovementService(apiClient)
  await queryClient.prefetchQuery(movementQueryOptions(movementService))

  return {
    token,
    dehydratedState: dehydrate(queryClient),
  }
}

export default function Home({ loaderData }: Route.ComponentProps) {
  return (
    <HydrationBoundary state={loaderData.dehydratedState}>
      <Movements />
    </HydrationBoundary>
  )
}
