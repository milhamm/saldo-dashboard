import type { z } from 'zod'
import { serverEnvSchema } from '~/schemas'

export function getServerEnv(): z.infer<typeof serverEnvSchema> {
  const parsedEnv = serverEnvSchema.safeParse(process.env)

  if (!parsedEnv.success) {
    const err = `Server environment variable validation failed: ${parsedEnv.error.message}`
    throw new Error(err)
  }

  return parsedEnv.data
}
