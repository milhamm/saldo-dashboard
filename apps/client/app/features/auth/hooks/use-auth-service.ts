import { useState } from 'react'
import { useApiClient } from '~/services/api'
import { AuthService } from '../service'

export function useAuthService() {
  const apiClient = useApiClient()
  const [authService] = useState(() => new AuthService(apiClient))
  return authService
}
