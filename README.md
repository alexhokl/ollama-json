# ollama-json

Ask question in natural language and answer the question in the specified JSON output

## Examples

Assuming schema file `sample_schema.json` with the following content.

```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "capital": {
      "type": "string"
    },
    "languages": {
      "type": "array",
      "items": {
        "type": "string"
      }
    }
  },
  "required": [
    "name",
    "capital",
    "languages"
  ]
}
```

Run the following command

```sh
ollama-json instruct -f sample_schema.json -q "Tell me about Canada"
```

