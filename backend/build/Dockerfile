# Шаг 1: Сборка приложения (используем многоступенчатую сборку)
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости (если нужны, например для CGO)
RUN apk add --no-cache git

# Копируем исходный код
COPY . .

# Компилируем приложение (указываем путь к main пакету)
RUN go build -C ./cmd/main/ -o ../../build/bin/app

# Шаг 2: Финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder-этапа
COPY --from=builder app/build/bin/app app

# Копируем .env
COPY --from=builder app/.env .

# Указываем точку входа
CMD ["app"]