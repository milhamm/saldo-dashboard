import { z } from 'zod'

export const serverEnvSchema = z.object({
  SESSION_SECRET: z.string(),
})

export const clientEnvSchema = z.object({
  VITE_API_BASE_URL: z.string(),
})
