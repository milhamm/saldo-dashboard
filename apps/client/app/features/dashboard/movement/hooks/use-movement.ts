import { queryOptions, useQuery } from '@tanstack/react-query'

import type { MovementService } from '../service'
import { useMovementService } from './use-movement-service'

export function movementQueryOptions(movementService: MovementService) {
  return queryOptions({
    queryKey: ['movements'],
    queryFn: () => movementService.getMovements(),
  })
}

export function useMovements(token: string) {
  const movementService = useMovementService(token)
  return useQuery(movementQueryOptions(movementService))
}
