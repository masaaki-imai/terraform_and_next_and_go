"use server";
import { cookies } from 'next/headers'

export async function customFetch<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const baseUrl = process.env.NEXT_PUBLIC_API_URL_CONTAINER

  const headers: Record<string, string> = {
    ...options.headers as Record<string, string>,
  }

  // FormDataの場合はContent-Typeを設定しない
  if (!(options.body instanceof FormData)) {
    headers['Content-Type'] = 'application/json'
  }

  const cookieStore = await cookies()
  const token = cookieStore.get('token')?.value
  if (token) {
    headers.Authorization = `Bearer ${token}`
  }

  const cookieHeader = cookieStore.getAll()
    .map(cookie => `${cookie.name}=${cookie.value}`)
    .join('; ')

  if (cookieHeader) {
    headers.Cookie = cookieHeader
  }

  const response = await fetch(`${baseUrl}${endpoint}`, {
    ...options,
    headers,
    credentials: 'include',
  })

  const setCookieHeader = response.headers.get('set-cookie');
  if (setCookieHeader) {
    const tokenMatch = setCookieHeader.match(/token=([^;]*)/);
    if (tokenMatch) {
      if (tokenMatch[1] === "") {
        cookieStore.delete('token');
      } else {
        cookieStore.set('token', tokenMatch[1], {
          path: '/',
          secure: process.env.NODE_ENV === 'production',
          sameSite: 'lax',
        });
      }
    }

    // line-account-sessionの処理を追加
    const lineAccountSessionMatch = setCookieHeader.match(/line-account-session=([^;]*)/);
    if (lineAccountSessionMatch) {
      if (lineAccountSessionMatch[1] === "") {
        cookieStore.delete('line-account-session');
      } else {
        cookieStore.set('line-account-session', lineAccountSessionMatch[1], {
          path: '/',
          secure: process.env.NODE_ENV === 'production',
          sameSite: 'lax',
          httpOnly: false,
          maxAge: 3600 * 24,
        });
      }
    }
  }

  if (!response.ok) {
    const errorData = await response.json();
    if (response.status === 422) {
      // バリデーションエラーの場合、詳細データを渡す
      return errorData;
    }
    console.error('API Error:', errorData)
    throw new Error(JSON.stringify(errorData))
  }

  return response.json();
}
