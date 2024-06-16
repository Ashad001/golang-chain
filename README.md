# My Go API

This is a Go-based API deployed on Hugging Face Spaces.

## How to Use

- Make sure you have set your system prompt in /prompts folder (by creating one)
    - Also, mention what json format you want your response to be, if you dont want in json, remove the line in `llms.WithJSONMode()` in `api/chain.go`
- Send a POST request to `/chat/` with the user query.
