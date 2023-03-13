# bankApp
bankApp - простая банковская система, позволяющая создавать аккаунты, создавать операции, проводить транзакции и отсылать/принимать деньги на счёт аккаунта. 

# Backend
## RESTful API
Приложение реализует RESTful архитектуру при помощи `Gin` - https://github.com/gin-gonic/gin.
Реализованы `GET`, `POST`, `UPDATE` и `DELETE` запросы.

## База данных и CRUD
Данные хранятся в базе данных `PostgreSQL`.

Миграция баз данных осуществлена при помощи golang-migrate - https://github.com/golang-migrate/migrate

CRUD операции позволяют
* Создать аккаунт, получить аккаунт или список аккаунтов, обновить аккаунт и его баланс, удалить аккаунт
* Создать операцию, получить информацию об операции или списке операций
* Создать трансфер, получить информацию о трансфере или списке трансферов

CRUD операции реализованы при помощи sqlc - https://github.com/kyleconroy/sqlc

Поддерживаются многопоточные запросы к БД

## Тестирование
Все тесты находятся в файлах с префиксом `_test` и реализованы при помощи библиотеки testify - https://github.com/stretchr/testify

Также в этом проекте я постарался имплементировать mock базу данных при помощи https://github.com/golang/mock и table-driven test подход. Тесты находятся в `api/account_test.go` файле.

## Конфигурация
Конфигурация реализована с помощью viper - https://github.com/spf13/viper. Настройки для конфигурации можно как задать в файле, так и в переменных среды при запуске сервера.

## Инструкция по сборке
Все необходимые инструкции находятся в Makefile. Последовательность:
1. `make postgres` - запуск docker контейнера с бд postgresql
2. `make createdb` - создание базы данных
3. `make migrateup` - наполнение базы данных
4. `make server` - запуск сервера на localhost
