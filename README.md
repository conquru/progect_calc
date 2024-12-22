# Go Calculator

Это простое приложение на Go, которое выполняет математические вычисления на основе предоставленных пользователем выражений. Приложение включает в себя API, который позволяет отправлять HTTP-запросы для вычисления математических выражений.

## Функциональность

- Поддержка базовых арифметических операций: сложение, вычитание, умножение и деление.
- Валидация входных выражений.
- Обработка ошибок, таких как недопустимые выражения или деление на ноль.

## Описание

Калькулятор принимает математические выражения в виде строки и возвращает результат. Он поддерживает следующие операции:
- Сложение (`+`)
- Вычитание (`-`)
- Умножение (`*`)
- Деление (`/`)
- Приоритет выполнения с использованием скобок

## Технологии

- Go (версия 1.16 и выше)
- Набор стандартных библиотек Go (`net/http`, `encoding/json` и др.)

## Установка

Для установки калькулятора на вашей машине выполните следующие шаги:

1. Убедитесь, что у вас установлен Go (версии 1.16 или выше). Вы можете скачать Go с [официального сайта](https://golang.org/dl/).
2. Клонируйте репозиторий:
```bash
   git clone https://github.com/conquru/project_calc.git
   cd project_calc
   go mod tidy
```
3. Запуск:
```bash
   go run .\cmd\main.go
```

## Таблица тестов

| request  | response |
| :---: | :---: |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "1+1"}'  | {"result": 2,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "(2+2)*2"}'  | {"result": 8,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "2+2*2"}'  | {"result": 6,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "3*3+4/2"}'  | {"result": 11,"error": "invalid"}  |
| curl --location 'localhost/api/v1/calculate' \--header 'Content-Type: application/json' \--data '{"expression": "4/2+4"}'  | {"error": "Expression is notvalid"}  |

