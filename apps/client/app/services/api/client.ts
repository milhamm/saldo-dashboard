import { type ConsolaInstance, createConsola } from 'consola'
import { type $Fetch, FetchError, type FetchRequest, ofetch } from 'ofetch'
import { getClientEnv } from '~/config'
import type { GenericError } from '~/types'

type ApiClientOptions = {
  baseUrl?: string
  bearerToken?: string
}

const isClientResponseError = (error: unknown): error is FetchError<GenericError> =>
  error instanceof FetchError && error.data

export class ApiClient {
  #options: ApiClientOptions
  #fetcher: $Fetch
  #logger: ConsolaInstance

  constructor() {
    this.#options = ApiClient.getDefaultOptions()
    this.#logger = createConsola({
      level: 4,
      defaults: {
        tag: 'ApiClient',
      },
    })
    this.#fetcher = this.#createFetcher(ofetch)
  }

  setOptions(opts: ApiClientOptions) {
    this.#options = opts
    return this
  }

  setBearerToken(token: string) {
    this.#options = {
      ...this.#options,
      bearerToken: token,
    }
    this.#fetcher = this.#createFetcher(ofetch)
    return this
  }

  #createFetcher(client: $Fetch) {
    const logger = this.#logger
    const opts = this.#options
    return client.create({
      baseURL: this.#options.baseUrl,
      onRequest(ctx) {
        logger.debug('[REQUEST]: ', ctx.request)
        if (opts.bearerToken) {
          ctx.options.headers.set('Authorization', `Bearer ${opts.bearerToken}`)
        }
      },
      onResponse(ctx) {
        logger.debug('[RESPONSE]: ', ctx.response)
      },
    })
  }

  async request<TData>(req: FetchRequest, opts: RequestInit) {
    const headers = new Headers(opts.headers)
    headers.append('Content-type', 'application/json')
    headers.append('Accept', 'application/json')

    try {
      const resp = await this.#fetcher<TData>(req, { ...opts, headers })
      return resp
    } catch (error) {
      if (isClientResponseError(error)) {
        error.message = error.data?.data || error.message
      }
      throw error
    }
  }

  static getDefaultOptions(): ApiClientOptions {
    return {
      baseUrl: getClientEnv().VITE_API_BASE_URL,
    }
  }
}
