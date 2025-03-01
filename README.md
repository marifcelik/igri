# İgri - Real-time Chat Application

A modern real-time chat application built with Go and React, featuring WebSocket communication and a clean, responsive UI.

## Features

- [ ] 🔐 Secure authentication system
- [X] 💬 Real-time messaging via WebSockets
- [X] 🌓 Light/Dark theme support
- [X] 📱 Responsive design
- [ ] 🌍 Interactive globe visualization
- [ ] 👑 Admin dashboard
- [X] 💾 Message history and conversation management

## Technology Stack

### Frontend
- React and Vite
- TanStack Router
- shadcn/ui

### Backend
- Go with Chi router
- MongoDB
- WebSocket server
- Redis for session management

## Development
1. Create a `.env` file and set `MONGO_PORT` and `REDIS_PORT` variables for port forwarding.
2. Just run `docker-compose up` to start the development environment.
