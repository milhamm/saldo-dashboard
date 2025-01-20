import type { GenericResponse } from '~/types'

export type MovementType = 'withdraw' | 'transfer' | 'top_up' | 'payment' | 'others'

export type Movement = {
  id: string
  amount: number
  fee: number
  movement_type: MovementType
  created_at: number
  updated_at: number
}

export type GetMovementResponse = GenericResponse<Movement[]>
