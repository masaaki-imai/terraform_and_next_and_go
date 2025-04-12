import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  experimental: {
    serverActions: {
      bodySizeLimit: '10mb'
    }
  },
  async redirects() {
    return [
      {
        source: '/',
        destination: '/public/org/login',
        permanent: false,
      },
    ]
  }
}

export default nextConfig
