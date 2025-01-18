import { ofetch } from 'ofetch'
import { Form, redirect } from 'react-router'
import { ApiClient } from '~/services/api'
import { createUserSession, getUserSession } from '~/services/auth.server'
import type { Route } from './+types/page'
import { AuthService } from './service'

export function meta() {
  return [{ title: 'Login' }, { name: 'description', content: 'Saldo Dashboard - Login' }]
}

export async function loader({ request }: Route.LoaderArgs) {
  const session = await getUserSession(request)
  if (session.has('access_token')) {
    return redirect('/')
  }
}

export async function action({ request }: Route.ActionArgs) {
  // const formData = await request.formData()
  let response: Response
  // const session = await getUserSession(request)

  const apiClient = new ApiClient(ofetch)
  const authService = new AuthService(apiClient)

  try {
    const token = await authService.login({
      phone: '082234',
      password: 'susudiamond1',
    })

    response = await createUserSession(request, token)
    if (!response) {
      throw new Error('An error occurred while creating the session')
    }
  } catch (error) {
    if (error instanceof Error) {
      return { error: error.message }
    }

    return { error: 'Unknown Error' }
  }
  return response
}

export default function Login({ actionData }: Route.ComponentProps) {
  return (
    <div>
      <Form method="post">
        <div className="capitalize">{actionData?.error}</div>
        <div>
          <div>
            <label htmlFor="phone">
              Phone:
              <input type="text" name="phone" id="phone" />
            </label>
          </div>
          <div>
            <label htmlFor="password">
              Password
              <input type="password" name="password" id="password" />
            </label>
          </div>
          <button type="submit">Login</button>
        </div>
      </Form>
    </div>
  )
}
