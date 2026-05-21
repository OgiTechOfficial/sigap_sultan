/** @type {import('next').NextConfig} */
// Configuration options for Next.js
// const basePath = process.env.NEXT_PUBLIC_BASE_PATH || '';

const nextConfig = {
  images: {
    domains: [
        'project.bi.sentech.id',
        'localhost'
    ],
  },
  //basePath: '',
  basePath: process.env.NODE_ENV === 'production' ? process.env.NEXT_PUBLIC_BASE_PATH : undefined,
  assetPrefix: process.env.NODE_ENV === 'production' ? process.env.NEXT_PUBLIC_BASE_PATH : undefined,
  reactStrictMode: true, // Enable React strict mode for improved error handling
  swcMinify: true,      // Enable SWC minification for improved performance
  compiler: {
    removeConsole: process.env.NODE_ENV !== "development", // Remove console.log in production
  },
};

console.log('Base Path:', process.env.NEXT_PUBLIC_BASE_PATH);
console.log('NODE_ENV:', process.env.NODE_ENV);

// Configuration object tells the next-pwa plugin 
const withPWA = require("next-pwa")({
  dest: "public", // Destination directory for the PWA files
  disable: process.env.NODE_ENV === "development", // Disable PWA in development mode
  register: true, // Register the PWA service worker
  skipWaiting: true, // Skip waiting for service worker activation
});

module.exports = withPWA(nextConfig);
