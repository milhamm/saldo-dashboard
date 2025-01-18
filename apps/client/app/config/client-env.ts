import type { z } from 'zod'
import { clientEnvSchema } from '~/schemas'

export function getClientEnv(): z.infer<typeof clientEnvSchema> {
  const parsedEnv = clientEnvSchema.safeParse(import.meta.env)

  if (!parsedEnv.success) {
    const err = `Client environment variable validation failed: ${parsedEnv.error.message}`
    throw new Error(err)
  }

  return parsedEnv.data
}
