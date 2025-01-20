import { useState } from 'react'
import { useApiClient } from '~/services/api'
import { MovementService } from '../service'

export function useMovementService(token: string) {
  const apiClient = useApiClient()
  apiClient.setBearerToken(token)
  const [authService] = useState(() => new MovementService(apiClient))
  return authService
}
