# REST API CRUD-operations

Пакет представляет собой планировщик задач (TODO-list), реализованный на REST API, PostgreSQL, Fiber. Позволяет:
- создавать задачу
- читать список задач
- обновлять статус задачи
- удалять задачу

API-эндпоинты:<br>
Post/tasks/add - создание задачи<br>
Get/tasks - получение списка всех задач<br>
Put/tasks/update - обновление статуса задачи<br>
Delete/tasks/delete/:taskIDToDelete - удаление задачи<br>

## Примеры выполнения CRUD-операций
###Получение списка всех имеющихся задач
[postman][img/img_1.jpg]
[postgres][img/img_2.jpg]

