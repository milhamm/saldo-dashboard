import { ApiClient } from '~/services/api'
import { createUserSession } from '~/services/auth.server'

import { loginRequestSchema as schema } from '../schema'
import { AuthService } from '../service'
import type { Route } from './+types/page'

export async function loginAction({ request }: Route.ActionArgs) {
  const formData = await request.formData()
  let response: Response

  const apiClient = new ApiClient()
  const authService = new AuthService(apiClient)

  const data = schema.safeParse(Object.fromEntries(formData.entries()))

  if (data.error) {
    return { error: data.error.message }
  }

  try {
    const token = await authService.login(data.data)
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
