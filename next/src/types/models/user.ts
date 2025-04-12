import { CommonFields } from './common'

export interface User extends CommonFields {
  id: number
  email: string
  password: string
  name: string
  company_name: string
  phone_number: number
  address_1: string
  address_2: string
  address_3: string
  post_code_1: number
  post_code_2: number
  last_time_login: Date
}