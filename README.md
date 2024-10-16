# README для вашего проекта

## Описание

Этот проект представляет собой простой HTTP-сервер, который обрабатывает события и взаимодействует с Kafka через конфигурационные файлы. Сервис написан на Go и использует Logrus для логирования.

## Структура проекта

```
event-router/
├── src/
│   ├── app/
│       ├── handlers/
│           └── event/
│               └── event_handler.go
│       ├── internal/
│           └── config/
│               ├── config.go
│               └── config_test.go
│   ├── main.go
├── go.mod
├── go.sum
```

## Установка и запуск

1. **Установите Go**: Убедитесь, что у вас установлена последняя версия Go.
2. **Клонируйте репозиторий**: 
    ```bash
    git clone https://github.com/ваш_никнейм/ваш_репозиторий.git
    cd ваш_репозиторий
    ```
3. **Установите зависимости**: 
    ```bash
    go mod tidy
    ```
4. **Создайте конфигурационный файл**: Создайте файл `config.yaml` в корне проекта со следующим содержимым:
    ```yaml
    service:
      host: localhost
      port: 8080
    kafka:
      host: localhost
      port: 9092
    ```
5. **Запустите сервер**:
    ```bash
    go run src/app/main.go
    ```

## Конфигурация

Конфигурация проекта загружается из файла `config.yaml`, который должен быть расположен в корне проекта. В случае отсутствия переменной окружения `CONFIG_PATH` будет выведено сообщение об ошибке, и сервер не запустится.

## Обработка событий

Сервер обрабатывает два маршрута:
- `/event`: Принимает POST-запросы для обработки событий.
- `/event/`: Принимает GET-запросы для получения информации о событии по его идентификатору.

## Тестирование

Для тестирования конфигурации используйте файл `config_test.go`. Временный файл с конфигурацией создается автоматически для каждого запуска теста, и удаляется после его окончания.

```bash
go test src/app/internal/config/config_test.go
```

## Логирование

В проекте используется Logrus для логирования. Уровень логирования установлен в `InfoLevel`. Дополнительные настройки логирования можно внести в файл конфигурации или переменных окружения.

## Зависимости

- **Logrus**: Для логирования.
- **YAML**: Для парсинга конфигурационных файлов.
- **Kafka-go**: (Комментарий) В случае, если вы планируете использовать Kafka для взаимодействия с событиями.

## Авторы

[Ваше Имя]

## Лицензия

Проект лицензирован под [выберите лицензию]. Подробнее в файле `LICENSE`.

---

Спасибо за использование и развитие проекта!