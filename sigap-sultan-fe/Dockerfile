FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package.json yarn.lock* ./

# Install dependencies
RUN yarn install --frozen-lockfile

# Copy all project files
COPY . .

# Expose Next.js dev server port
EXPOSE 3000

# Run the app in development mode
CMD ["yarn", "dev", "-p", "3000", "-H", "0.0.0.0"]
