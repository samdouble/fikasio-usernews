# fikasio-usernews
API to query rideshares posted on AmigoExpress / Kangaride

## Development
```
./app dev
```

## Production
Just push on any branch and CircleCI will deploy on AWS ECR and AWS Lambda.

### Environment variables

| Name          | Required  | Default  |   |   |
|---------------|-----------|----------|---|---|
| MONGODB_URL   | Yes       |          |   |   |
