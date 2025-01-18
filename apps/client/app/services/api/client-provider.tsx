import { createContext, use } from 'react'
import type { ApiClient } from './client'

const ApiClientContext = createContext<ApiClient | null>(null)

export function ApiClientProvider({
  children,
  client,
}: React.PropsWithChildren<{ client: ApiClient }>) {
  return <ApiClientContext.Provider value={client}>{children}</ApiClientContext.Provider>
}

export function useApiClient() {
  const client = use(ApiClientContext)

  if (!client) {
    throw new Error('useApiClient must be used inside of ApiClientProvider')
  }

  return client
}
