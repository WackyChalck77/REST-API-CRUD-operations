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
### Получение списка всех имеющихся задач
Используется метод GET http://127.0.0.1:3000/tasks/.<br>
Происходит выдача таблицы задач,
[Postman](img/img_1.jpg).
Таблица в Postgres,
[Postgres](img/img_2.jpg).

### Создание новой задачи
Используется метод POST http://127.0.0.1:3000/tasks/add/. Данные передаются в следующем виде (JSON): <br>
{  
"Title": "Новая тестовая задача",  
"Description": "Описание новой тестовой задачи",  
"Status": "new"  
}  
Примечание: поле Status может принимать только следующие значения: 'new', 'in_progress', 'done'.

Происходит вставка задачи в таблицу,
[Postman](img/img_3.jpg).
Таблица в Postgres,
[Postgres](img/img_4.jpg).


