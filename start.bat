@echo off

rem Запрос сообщения у пользователя
set /P "message=Wire expr: "

rem Сложение сообщения с строкой
set "full_message=%message%"

rem Отправка запроса
curl -X POST -d "expression=%full_message%" http://localhost:8080/expression

curl -X GET -o response.json http://localhost:8080/expressions

rem Прочитать файл response.json
for /f "delims=" %%i in ('type response.json') do (
  echo %%i
)

rem Пауза для просмотра ответа
pause
