# toy-ses

Simple Email Service

## 🚧 Under Development 🚧

## Start server

```bash
GIN_MODE=release go run .
```

## Send Email

```
POST /api/email/send/<emai>

Header:
    Content-Type: application/json

Body: {"subject":"<subject>","body":"<body>"}
```

#### Response On Success

```json
{"status":"success"}
```

#### Response On Failure

```json
{"status":"faild"}
```

## Count Email

```
GET /api/email/count
```

#### Response

```json
{"faild_count":<number>,"success_count":<number>}
```

## Test

```
GIN_MODE=release go test
```
