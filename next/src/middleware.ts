import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

// 認証タイプの定義
type AuthType = 'org' | 'admin'

// ディレクトリと認証タイプのマッピング
const directoryAuth: Record<string, AuthType> = {
  'org': 'org',
  'admin': 'admin',
}

export async function middleware(request: NextRequest) {
  console.log('Middleware executing for path:', request.nextUrl.pathname);
  const path = request.nextUrl.pathname

  // publicディレクトリのパスは認証不要
  if (path.startsWith('/public')) {
    return NextResponse.next()
  }

  // ディレクトリ名を取得（最初のパスセグメント）
  const directory = path.split('/')[1]

  // ディレクトリに対応する認証タイプを取得
  const requiredAuthType = directoryAuth[directory]

  // 認証タイプが未定義の場合は404を返す
  // （全てのパスは /public, /org, /admin のいずれかに属するべき）
  if (!requiredAuthType) {
    return new NextResponse(null, { status: 404 })
  }

  const token = request.cookies.get('token')

  // トークンがない場合は対応するログインページへ
  if (!token) {
    return redirectToLogin(request, directory)
  }
}

// ログインページへのリダイレクト
function redirectToLogin(request: NextRequest, directory: string) {
  const url = request.nextUrl.clone()
  // ログインページはpublicディレクトリ配下に配置
  url.pathname = `/public/${directory}/login`
  url.search = `?redirectTo=${encodeURIComponent(request.nextUrl.pathname)}`
  return NextResponse.redirect(url)
}

// ミドルウェアを適用するパスの設定
export const config = {
  matcher: [
    '/org/:path*',
    '/admin/:path*',
    '/public/:path*',
  ]
}
