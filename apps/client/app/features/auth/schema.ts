import { z } from 'zod'

const phone = z.string()
const password = z.string().min(1).min(8)
const token = z.string()

export const loginRequestSchema = z.object({
  phone,
  password,
})

export const loginResponseSchema = z.object({
  token,
})

export type LoginRequest = z.infer<typeof loginRequestSchema>
export type LoginResponse = z.infer<typeof loginResponseSchema>
