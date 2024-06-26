# Car Rental Service

## Описание

Этот сервис позволяет бронировать автомобили, подсчитывать стоимость аренды и формировать отчеты о загрузке автомобилей.

## Требования

- Go 1.20 или выше
- Docker
- PostgreSQL

## Конфигурация

Конфигурация сервиса разделена на два файла:

1. `config.yaml` - используется для хранения параметров конфигурации.
2. `.env` - используется для хранения конфиденциальной информации, такой как пароль к базе данных.

### config.yaml
Измените файл `config.yaml` в `configs` со следующим содержимым:

```yaml
port: "8080"
db:
  host: "localhost"
  username: "postgres"
  name: "postgres"
  ```

`.env`

Создайте файл `.env` в корне проекта со следующим содержимым:

```DB_PASSWORD=your_db_password```

## Запуск

### Локальный запуск
Убедитесь, что у вас запущен сервер PostgreSQL и доступен с параметрами, указанными в config.yaml и .env.

Установите зависимости и запустите приложение:
```
go mod download
go run cmd/main.go
```
### Запуск с использованием Docker
Постройте Docker образ:
```docker build -t car-rental-app .```

Запустите контейнер:
``` docker run -p 8080:8080 --env-file=.env car-rental-app```

-----------
Предполагается, что запущен сервер Postgres, в базе данных созданы необходимые таблицы и добавлены 5 автомобилей в таблицу cars.

Для запуска приложения и Postgres в одной сети:

1. Создайте сеть Docker:
```bash
docker network create car-rental-net
```

2. Запустите контейнер PostgreSQL:
```bash
docker run --name car-rental-postgres --network car-rental-net -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=car_rental -p 5432:5432 -d postgres
```
Создайте необходимые таблицы (схема в файле bd_cheme) и добавьте записи о 5 машинах в таблицу cars

3. Запустите контейнер с приложением, подключив его к той же сети:
```bash
docker run --name car-rental-app --network car-rental-net -p 8080:8080 --env-file=.env car-rental-app
```