import { ApiClient } from '~/services/api/client'
import type { GetMovementResponse } from './types'

export class MovementService {
  #client: ApiClient

  constructor(client: ApiClient) {
    this.#client = client
  }

  async getMovements() {
    const res = await this.#client.request<GetMovementResponse>('/movements', { method: 'GET' })
    return res.data
  }
}
