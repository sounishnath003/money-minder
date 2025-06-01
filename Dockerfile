# -- Build the backend application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy && go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-s -w" -o /app/bin/moneyminder cmd/*.go

# -- Build the frontend application
FROM node:22-alpine AS frontend-builder

WORKDIR /app

COPY web/package.json .
COPY web/package-lock.json .
RUN npm install

COPY web/ .

# Add build argument for commit hash
ARG VITE_COMMIT_HASH
ENV VITE_COMMIT_HASH=$VITE_COMMIT_HASH

RUN npm run build

# -- And finally, build the final image
FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/bin/moneyminder /app/bin/moneyminder
COPY --from=frontend-builder /app/dist/ /app/web/dist/

RUN chown -R 1000:1000 /app

USER 1000

EXPOSE 3000

ENTRYPOINT [ "/app/bin/moneyminder" ]