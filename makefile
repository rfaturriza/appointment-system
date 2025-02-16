.PHONY: run-backend run-frontend run-all backend-build frontend-build

# Backend commands
run-backend:
	cd backend && go run main.go

backend-build:
	cd backend && go build -o app

# Frontend commands
run-frontend:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build

# Run both applications
run-all:
	make run-backend & make run-frontend

# Install dependencies
install:
	cd backend && go mod tidy
	cd frontend && npm install

# Clean
clean:
	cd backend && go clean
	cd frontend && rm -rf node_modules .nuxt dist