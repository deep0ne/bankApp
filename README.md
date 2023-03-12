# bankApp
bankApp - простое банковское приложение, позволяющее создавать аккаунты, проводить транзакции и отсылать/принимать деньги на счёт.

# Backend
1. Приложение реализует RESTful архитектуру при помощи Gin - https://github.com/gin-gonic/gin.
2. Данные об аккаунтах и юзерах хранятся в базе данных *PostgreSQL*. Миграция баз данных осуществлена при помощи golang-migrate - https://github.com/golang-migrate/migrate
3. CRUD операции реализованы при помощи sqlc - https://github.com/kyleconroy/sqlc
