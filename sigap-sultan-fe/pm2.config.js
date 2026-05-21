module.exports = {
  apps: [
    {
      name: 'next-app',
      script: 'node_modules/.bin/next',
      args: 'start',
      cwd: '/opt/www/bi/fe/', // Path to your project root
      instances: 1, // Number of instances
      autorestart: true,
      watch: false,
      max_memory_restart: '2G',
      env: {
        NODE_ENV: 'production',  // Ensure it's running in production mode
        NEXT_PUBLIC_BASE_PATH: '/web/v1/stg',  // Optionally, set environment variables directly
	DEBUG: '*', // Enable verbose logging for debug
      },
    },
  ],
};

