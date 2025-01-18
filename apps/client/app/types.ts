export type GenericResponse<T> = {
  code: number
  message: string
  data: T
}

export type GenericError = GenericResponse<string>
