const isProd = process.env.NODE_ENV === 'production';

export const assetPrefix = (url: string) => {
  if (!isProd) {
    return url;
  }
  return process.env.NEXT_PUBLIC_BASE_PATH + url;
}