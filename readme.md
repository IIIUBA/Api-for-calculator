Система LMS (Lyceum Math System) 
Вычисление арифметических выражений, используя веб-интерфейс или API.

## Особенности:

- Метод: POST
- Веб-интерфейс: -
- Масштабируемость: Возможность распределения вычислений по агента

## Установка
1.
```
git clone https://github.com/IIIUBA/lms.git
```
```
Run Code ( Cltr + Alt + N ) или go run main.go
```

## API:

```
curl -X POST -d "expression=2p3*4" http://localhost:8080/expression
```
 - Отправить выражение, где " 2p3*4 " ваш пример | в http ельзя использовать +, поэтому вместо него пишите - p

```
curl -X GET http://localhost:8080/expressions
``` 
- Статус выражений.

```
curl -X GET http://localhost:8080/agents_status
``` 
- Статус агентов.

```
curl -X POST -d "add=2" http://localhost:8080/computation_agent
``` 
- Добавить агента(ов), где " add=2 " кол-во агентов


## Вручную:

Статус агентов
```
http://localhost:8080/agents_status
```
Статус выражений
```
http://localhost:8080/expressions
```
## Дополнительно

Запустите start.bat и введите пример
```
Ввести можно будет только пример и получить ответ
```
