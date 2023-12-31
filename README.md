# url-shortener
## Укоротитель ссылок. Реализация тестового задания AltCraft
ТЗ:

    https://github.com/altkraft/for-applicants/blob/master/backend/shortener/task.md
## Используемые технологии
    github.com/sirupsen/logrus
    github.com/spf13/viper
    github.com/redis/go-redis/v9
    github.com/gin-gonic/gin
    github.com/swaggo/swag/cmd/swag
    github.com/swaggo/gin-swagger
## Database - Redis with persistent storage
    docker run --name url-shortener-redis-db -p 6001:6379 -d redis redis-server --save 60 1 --loglevel warning
## Количество хранимых ключей
Redis может поддерживать хранение 2^32 ключей

    https://redis.io/docs/getting-started/faq/#:~:text=Redis%20can%20handle%20up%20to,available%20memory%20in%20your%20system.

## Куда должно развиваться приложение
Приложение должно 
* иметь удобный и минималистичный графический интерфейс
* иметь более устойчивую к коллизиям систему хеширования url
* иметь возможность для сбора статистики
* очистка для неиспользуемых url-хешей
* полноценное развертывание в docker