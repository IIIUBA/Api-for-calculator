Система LMS (Lyceum Math System) вычисление арифметических выражений, используя веб-интерфейс или API.

** Особенности:**

Веб-интерфейс: HTML + Vue.js
API: RESTful API
База данных: MySQL
Масштабируемость: Возможность распределения вычислений по агентам
Тестирование: Тестовые выражения в базе данных
** Запуск:**

Установите Docker и Docker Compose.
git clone https://github.com/iiiuba/lms.git
***Start***
```
docker-compose up -d --build
```
***Stop***
```
docker-compose down
```
Откройте веб-интерфейс: http://127.0.0.1/
** API:**

GET /api/expressions - список выражений.
GET /api/expressions/2 - конкретное выражение.
POST /api/expressions - добавить выражение.
GET /api/get-new-task - получить новое выражение.
GET /api/settings - список настроек.
POST /api/settings - сохранить настройки.
GET /api/workers - список агентов.
** Дополнительно:**

Подключение к базе данных:
```
Хост: 127.0.0.1
Порт: 3306
```
```
Пользователь: root
Пароль: testerum
```
** Ссылки:**

Яндекс Лицей: https://yandexlyceum.ru/
Docker: https://www.docker.com/
Docker Compose: https://docs.docker.com/compose/
Vue.js: https://vuejs.org/
MySQL: https://www.mysql.com/