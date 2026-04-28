# remoteok_vue_go

Project structure:

```text
.
├── frontend/  # Vue 3 + Vite app
└── backend/   # Go API
```

## Frontend

```sh
cd frontend
npm install
npm run dev
```

Build:

```sh
cd frontend
npm run build
```

## Backend

Run the Go server:

```sh
cd backend
go run .
```

The backend starts on `http://localhost:8080` and exposes `GET /health`.
