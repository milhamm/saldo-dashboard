import { z } from 'zod'

const phone = z.string()
const password = z
  .string()
  .min(1, { message: 'Kata sandi harus diisi' })
  .min(8, { message: 'Kata sandi minimal 8 karakter' })
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
