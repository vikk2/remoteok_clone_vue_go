# remoteok_vue_go

Project structure:

```text
.
├── frontend/  # Vue 2 + Vite app
└── backend/   # Go API
```

## Frontend

```sh
cd frontend
npm install
npm run dev
```

## Backend

Run the Go server:

```sh
cd backend
go run main.go
```

The backend starts on `http://localhost:3000` and exposes `GET /health`.
