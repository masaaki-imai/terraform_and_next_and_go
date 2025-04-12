export interface CommonFields {
  ins_timestamp: string
  ins_user_id: number
  ins_action: string
  upd_timestamp: string
  upd_user_id: number
  upd_action: string
  del_timestamp: string | null
  del_user_id: number | null
  del_action: string | null
}
