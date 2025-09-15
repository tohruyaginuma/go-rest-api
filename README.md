# Hands-on GO and Echo from Udemy

## URL

https://www.udemy.com/course/echo-go-react-restapi/learn/lecture/37074464#content

## Migrations

```
GO_ENV=dev go run migrate/migrate.go
```

## Login to database

```bash
docker exec -it go-rest-api-dev-postgres-1 psql -U udemy -d udemy
```
