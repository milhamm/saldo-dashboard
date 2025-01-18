import { ApiClient } from '~/services/api/client'
import type { LoginRequest, LoginResponse } from './schema'

export class AuthService {
  #client: ApiClient
  static instance: AuthService

  constructor(client: ApiClient) {
    this.#client = client
  }

  async login(payload: LoginRequest) {
    const res = await this.#client.request<LoginResponse>('/auth/login', {
      method: 'POST',
      body: JSON.stringify(payload),
    })
    return res.token
  }
}
