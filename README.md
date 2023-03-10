# Asteroids NeoWs API

API для сервиса NASA Asteroids NeoWs, который позволяет получить количество астероидов, находящихся рядом с Землей, за
определённую дату.

## Как работать с API

Сервис имеет REST-архитектуру со следующими методами:

### `GET`

Возвращает количество астероидов по дате. На вход приходят разные даты в формате `ҮҮҮҮ-ММ-DD`:

```
/neo/count?dates=2020-01-20&dates=2020-03-20&dates=2020-05-24&dates=2020-08-22
```

### `POST`

Записывает (или обновляет, если такая дата уже есть) данные в БД.
На вход по адресу `/neo/count` приходит массив из дат и количества астероидов в JSON:

```
{
  "neo_counts": [
    {
      "date": "2020-01-20",
      "count": 12
    },
    {
      "date": "2020-02-26",
      "count": 9
    }
  ]
}
```
