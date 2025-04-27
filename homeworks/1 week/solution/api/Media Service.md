
```yaml
name: Media Service
responsibility: Хранение и выдача аватарок/файлов

data: Files (S3, etc)
api:
  - POST /upload
  - GET /file
publishes_events:
  - FileUploaded
consumes_events: []
interacts_with:
  - User/Profile Service
```

