import { redirect } from 'react-router'
import { commitSession, destroySession, sessionStorage } from './session.server'

export async function getUserSession(req: Request) {
  return await sessionStorage.getSession(req.headers.get('Cookie'))
}

export async function destroyUserSession(req: Request) {
  const session = await getUserSession(req)
  return redirect('/login', {
    headers: {
      'Set-Cookie': await destroySession(session),
    },
  })
}

export async function createUserSession(req: Request, token: string) {
  const session = await getUserSession(req)
  session.set('access_token', token)
  return redirect('/', {
    headers: {
      'Set-Cookie': await commitSession(session, {
        httpOnly: true,
        sameSite: 'lax',
        maxAge: 60 * 60 * 24 * 3,
      }),
    },
  })
}
