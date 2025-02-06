# fikasio-usernews

## Development

### Running locally

```
./app dev
```

### Installing dependencies

```
./app enter
go get <PACKAGENAME>
```

## Production
Just push on any branch and CircleCI will deploy on AWS ECR and AWS Lambda.

### Environment variables

| Name            | Required  | Default  |   |   |
|-----------------|-----------|----------|---|---|
| MONGODB_URL     | Yes       |          |   |   |
| SMTP_HOST       | Yes       |          |   |   |
| SMTP_PORT       | Yes       |          |   |   |
| SMTP_USERNAME   | Yes       |          |   |   |
| SMTP_PASSWORD   | Yes       |          |   |   |
