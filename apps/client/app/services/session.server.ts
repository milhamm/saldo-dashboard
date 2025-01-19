import { createCookieSessionStorage } from 'react-router'
import { getServerEnv } from '~/config'

type SessionData = {
  access_token: string
}

export const sessionStorage = createCookieSessionStorage<SessionData>({
  cookie: {
    name: '__session',
    sameSite: 'lax',
    path: '/',
    httpOnly: true,
    secrets: [getServerEnv().SESSION_SECRET],
    secure: process.env.NODE_ENV === 'production',
  },
})

export const { commitSession, destroySession, getSession } = sessionStorage
