import { createCookieSessionStorage } from 'react-router'
import { getServerEnv } from '~/config'

export const sessionStorage = createCookieSessionStorage({
  cookie: {
    name: '__session',
    sameSite: 'lax',
    path: '/',
    httpOnly: true,
    secrets: [getServerEnv().SESSION_SECRET],
    secure: process.env.NODE_ENV === 'production',
  },
})

export const { commitSession, getSession } = sessionStorage
