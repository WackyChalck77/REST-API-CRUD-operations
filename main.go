package main

import (
	"Go/06_rest/model"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// глобальные переменные для подключения к БД
var (
	driverName = "postgres"
	dataSource = "host=localhost user=postgres password=21082014 dbname=rest_api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func main() {

	//установка времени
	var timeNowString string
	//время для Москвы
	moscowLocation, _ := time.LoadLocation("Europe/Moscow")
	//текущее время
	now := time.Now().In(moscowLocation)
	//форматирование времени под формат Postgres
	timeFormated := now.Format("2006-01-02 15:04:05.99")
	//преобразуем в строку
	timeNowString = timeFormated
	//timeNow := fmt.Sprint(timeFormated)
	fmt.Println(timeNowString)

	//Struct Для внесения новой задачи
	//в JSON
	// {
	// 	"Title":       "Тестовая задача",
	// 	"Description": "Описание тестовой задачи",
	// 	"Status":      "new",
	// 	}

	//Struct для обновления статуса задачи под определенным номером
	//в JSON
	// {
	// 	"ID": 21,
	// 	"Status": "done"
	// 	}

	app := fiber.New()
	//прямая ссылка
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Добро пожаловать на REST API.")
		if err != nil {
			return err
		}
		return nil
	})

	//Показ всех задач API
	app.Get("/tasks", GetAllTasks)

	//Добавление задачи API
	app.Post("/tasks/add", PostTask)

	//Удаление задачи API
	app.Delete("/tasks/delete/:taskIDToDelete", DeleteTaskByID)

	//Обновление статуса задачи API
	app.Put("/tasks/update", UpdateTask)

	//работа сервера
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

// функция соединения с БД
func Connect() *sql.DB {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer db.Close()
	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение к базе данных")
	return db
}

// Читать список задач
func SelectAlltasks() []model.Task {
	db := Connect()
	rows, err := db.Query("SELECT * from tasks")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	//пустой срез, хранит результаты выборки
	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			log.Fatal(err.Error())
		}
		tasks = append(tasks, task)
	}
	return tasks
}

// функция для Fiber показать список задач
func GetAllTasks(c *fiber.Ctx) error {
	err := c.SendString("Пытаемся достать все строки таблицы")
	if err != nil {
		return err
	}
	allTasks := SelectAlltasks()
	return c.JSON(allTasks)
}

// Создать задачу
func CreateTask(test_task model.Task) string {
	db := Connect()
	defer db.Close()
	create, err := db.Prepare("insert into tasks (title, description, status) values ($1, $2, $3);")
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := create.Exec(test_task.Title, test_task.Description, test_task.Status)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
	var lastInsertedID int
	err = db.QueryRow("select max(id)from tasks").Scan(&lastInsertedID)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(lastInsertedID)
	lastInsertedIDString := fmt.Sprint(lastInsertedID)
	fmt.Printf("\n Ззадача [%v] успешно добавлена под номером [%v]", test_task.Title, lastInsertedIDString)
	returnString := fmt.Sprintf("\n Ззадача [%v] успешно добавлена под номером [%v]", test_task.Title, lastInsertedIDString)
	return returnString
}

// функция для Fiber Создавать задачу
func PostTask(c *fiber.Ctx) error {
	var task model.Task
	err := c.BodyParser(&task)
	if err != nil {
		return err
	}
	postedTaskID := CreateTask(task)
	//возвращает id добавленной задачи
	return c.SendString(postedTaskID)
}

// Обновить задачу (статус)
func UpdateTaskStatus(update_task model.Task) interface{} {
	db := Connect()
	defer db.Close()
	update, err := db.Prepare("update tasks set status=$1, updated_at=NOW() where id=$2")
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := update.Exec(update_task.Status, update_task.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
	fmt.Println("Изменен статус задачи", update_task.ID)
	return update_task.ID
}

// функция для Fiber обновить задачу
func UpdateTask(c *fiber.Ctx) error {
	var task model.Task
	err := c.BodyParser(&task)
	if err != nil {
		return err
	}
	rowUpdated := UpdateTaskStatus(task)
	return c.SendString(fmt.Sprintf("Обновлена задача %v", rowUpdated))
}

// Удалить задачу
func DeleteTask(id int) interface{} {
	db := Connect()
	defer db.Close()
	delete, err := db.Prepare("delete from tasks where id=$1")
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := delete.Exec(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}
	return rowsAffected
}

// функция для Fiber удалить задачу
func DeleteTaskByID(c *fiber.Ctx) error {
	taskIDToDelete, err := strconv.Atoi(c.Params("taskIDToDelete"))
	if err != nil {
		log.Fatal(err.Error())
	}
	DeleteTask(taskIDToDelete)
	return c.SendString("Удалена задача ID: " + strconv.Itoa(taskIDToDelete))
}
