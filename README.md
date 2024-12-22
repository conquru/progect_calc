# Go Calculator

Это простое приложение на Go, которое выполняет математические вычисления на основе предоставленных пользователем выражений. Приложение включает в себя API, который позволяет отправлять HTTP-запросы для вычисления математических выражений.

## Функциональность

- Поддержка базовых арифметических операций: сложение, вычитание, умножение и деление.
- Валидация входных выражений.
- Обработка ошибок, таких как недопустимые выражения или деление на ноль.

## Технологии

- Go (версия 1.16 и выше)
- Набор стандартных библиотек Go (`net/http`, `encoding/json` и др.)

## Таблица тестов

| request  | response |
| :---: | :---: |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "1+1"}'  | {"result": 2,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "(2+2)*2"}'  | {"result": 8,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "2+2*2"}'  | {"result": 6,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "3*3+4/2"}'  | {"result": 11,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "4/2+4"}'  | {"error": "Expression is notvalid"}  |
