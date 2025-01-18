import { redirect } from 'react-router'
import { commitSession, sessionStorage } from './session.server'

export async function getUserSession(req: Request) {
  return await sessionStorage.getSession(req.headers.get('Cookie'))
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
