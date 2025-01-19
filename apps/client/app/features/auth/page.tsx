import { zodResolver } from '@hookform/resolvers/zod'
import {
  Alert,
  AlertDescription,
  AlertTitle,
  Button,
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Input,
} from '@saldo-dashboard/shared-ui'
import { AlertCircle } from 'lucide-react'
import { ofetch } from 'ofetch'
import { useForm } from 'react-hook-form'
import { redirect, useFetcher } from 'react-router'
import { ApiClient } from '~/services/api'
import { createUserSession, getUserSession } from '~/services/auth.server'
import type { Route } from './+types/page'

import { loginRequestSchema } from './schema'
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
  const formData = await request.formData()
  let response: Response

  const apiClient = new ApiClient(ofetch)
  const authService = new AuthService(apiClient)

  const data = loginRequestSchema.safeParse(Object.fromEntries(formData.entries()))

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

export default function Login(_: Route.ComponentProps) {
  const fetcher = useFetcher<typeof action>()
  const isBusy = fetcher.state !== 'idle'
  const loginForm = useForm({
    resolver: zodResolver(loginRequestSchema),
  })

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <Card>
            <CardHeader>
              <CardTitle className="text-2xl">Masuk</CardTitle>
            </CardHeader>
            <CardContent>
              <Form {...loginForm}>
                <form
                  onSubmit={loginForm.handleSubmit(async (data) => {
                    await fetcher.submit(data, { action: '/login', method: 'POST' })
                  })}
                >
                  <div className="flex flex-col gap-6">
                    <div className="grid gap-2">
                      <FormField
                        disabled={isBusy}
                        control={loginForm.control}
                        name="phone"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Nomor HP</FormLabel>
                            <FormControl>
                              <Input placeholder="Masukkan Nomor HP" {...field} />
                            </FormControl>
                            <FormDescription>e.g: 08123456789</FormDescription>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                    <div className="grid gap-2">
                      <FormField
                        disabled={isBusy}
                        control={loginForm.control}
                        name="password"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Kata Sandi</FormLabel>
                            <FormControl>
                              <Input type="password" placeholder="Masukkan kata sandi" {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                    {fetcher.data?.error ? (
                      <Alert variant="destructive">
                        <AlertCircle className="h-4 w-4" />
                        <AlertTitle>Error</AlertTitle>
                        <AlertDescription className="capitalize">
                          {fetcher.data?.error}
                        </AlertDescription>
                      </Alert>
                    ) : null}
                    <Button disabled={isBusy} type="submit" className="w-full">
                      {isBusy ? '...' : 'Masuk'}
                    </Button>
                  </div>
                  {/* <div className="mt-4 text-center text-sm">
                  Don&apos;t have an account?{' '}
                  <a href="#" className="underline underline-offset-4">
                    Sign up
                  </a>
                </div> */}
                </form>
              </Form>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
