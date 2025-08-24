# About app

App to fetch request to API to get definition of vocabs

# Run app locally

## Environment variables

See `.env`

Run

```bash
go run cmd/main.go
```

# AI API

## ChatGPT

- Curl
  ```
  curl "https://api.openai.com/v1/responses" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d '{
      "model": "gpt-5",
      "input": "Write a one-sentence bedtime story about a unicorn."
  }'
  ```

## Gemini

TODO
