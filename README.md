<!-- ---
title: My Go API
emoji: 🐹
colorFrom: blue
colorTo: green
sdk: docker
sdk_version: "latest"
app_file: .huggingface/Dockerfile
pinned: false
--- -->

# Go LangChain API

This is a Go-based API deployed on Hugging Face Spaces.

## How to Use

- Send a POST request to `/chat/` with the user query.

### Example System Prompt
```json
{
  "prompt": "
  Your answer should only be in JSON format with this format instruction and nothing else:
  {
    \"score\": <int>, 
    \"description\": \"<string>\",
    \"messages\": [
      {
        \"role\": \"<string>\",
        \"content\": \"<string>\"
      }
    ]
  }
}
```
**NOTE**: You have to make sure user query is related to the prompt too for good results
s