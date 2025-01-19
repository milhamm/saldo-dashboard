import { createContext, use } from 'react'

const AuthContext = createContext<string | null>(null)

export function AuthProvider({ token, children }: React.PropsWithChildren<{ token: string }>) {
  return <AuthContext.Provider value={token}>{children}</AuthContext.Provider>
}

export function useAuth() {
  const token = use(AuthContext)
  if (!token) {
    throw new Error('useAuth must be used inside of AuthProvider')
  }
  return token
}
