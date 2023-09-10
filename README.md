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