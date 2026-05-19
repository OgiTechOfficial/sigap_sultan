FROM node:18-alpine

WORKDIR /app

# Copy package files
COPY package.json package-lock.json* ./

# Install dependencies
RUN npm install

# Copy all project files
COPY . .

# Expose the default angular dev server port
EXPOSE 4200

# Run the app
CMD ["npm", "start", "--", "--host", "0.0.0.0"]
