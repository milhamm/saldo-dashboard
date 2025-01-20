import { type UIMatch, useMatches } from 'react-router'

type Data =
  | {
      token?: string
    }
  | undefined

export function useAccessToken() {
  const matches = useMatches() as UIMatch<Data>[]
  const match = matches.find((m) => m.data?.token)
  const token = match?.data?.token
  if (!token) {
    throw new Error('Token not found: You might forgot to return token from `loader`')
  }
  return token
}
