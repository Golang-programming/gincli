version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: "{{.AppName}}_app"
    environment:
      - DB_HOST=db
      - DB_USER=${DB_USER}
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=${DB_NAME}
      - DB_PORT=3306
    ports:
      - "8080:8080"
    depends_on:
      - db

  db:
    image: mysql:5.7
    container_name: "{{.AppName}}_db"
    environment:
      MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}
      MYSQL_DATABASE: ${DB_NAME}
      MYSQL_USER: ${DB_USER}
      MYSQL_PASSWORD: ${DB_PASSWORD}
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

volumes:
  db_data:
