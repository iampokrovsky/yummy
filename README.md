# Задание

1. Используя сущности, описанные на 5 недели (или написать новые) необходимо подготовить http методы, использующие
   методы
   репозитория;
2. Покрыть юнит-тестами ручки из п.1. Минимальное покрытие - 40%. Допускается использование композиций и тестирование
   их;
3. Покрыть интеграционными тестами методы репозитория дз 5 недели. Минимальные тест кейсы - успешное выполнение,
   получение
   ошибки из-за передачи некорректных данных;
4. Подготовить Makefile, в котором будут следующие команды: запуск тестового окружения при помощи docker-compose, запуск
   интеграционных тестов, запуск юнит-тестов, запуск скрипта миграций, очищение базы от тестовых данных;
5. (💎) Покрыть интеграционными тестами ручки из п.1.

# Реализация

Покрыт интеграционными тестами репозиторий `Меню`, для него написана HTTP-хэндлеры, которые также покрыты юнит и
интеграционными тестами.

## Как запустить

1. Переименовать файл `.env.sample` в `.env`;
2. Поднять окружение командой `make compose-up`;
3. Выполнить миграции командой `make migrate-up`;
4. Запустить юнит-тесты командой `make test-unit`;
5. Запустить интеграционные тесты командой `make test-integration`;
6. Проверить покрытие тестами командой `make test-cover`. Отчет появится в корне проекта в файле `cover.html`.
