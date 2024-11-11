FROM node:22.1.0-alpine as builder
WORKDIR /app
RUN npm install -g pnpm
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY . .
RUN pnpm build

FROM caddy:2.7.6-alpine
COPY --from=builder /app/dist /var/www
COPY Caddyfile /etc/caddy/Caddyfile

EXPOSE 80