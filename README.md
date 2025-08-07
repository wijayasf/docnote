# docNote

docNote is a lightweight Go library for Document Notarization Chain, a minimal proof-of-work blockchain recording document hashes with timestamps and providing a simple HTTP API.

## Features

- Go package `notary` with functions:
  - `Notarize(data string) Block`
  - `GetChain() []Block`
- HTTP server with endpoints:
  - `POST /notarize` — submit a document hash
  - `GET  /chain`    — retrieve all blocks
- Configurable proof-of-work difficulty (default: 3)
- Unit tests for blockchain logic
- Easy integration with automation tools (n8n, Postman, CLI)

## Installation

Get the module via:
```bash
go get github.com/wijayasf/docnote@v0.1.0
```

Or clone the repository:
```bash
git clone git@github.com:wijayasf/docnote.git
cd docnote
```

## Usage

### As a Go library
```go
import "github.com/wijayasf/docnote/notary"

func main() {
    block := notary.Notarize("yourDocumentHash")
    fmt.Println("New block:", block)
    chain := notary.GetChain()
    fmt.Println("Blockchain:", chain)
}
```

### Running the HTTP server
```bash
go run cmd/server/main.go
# Server listens on http://localhost:8080
```

- **POST /notarize**
  - Request body (JSON):
    ```json
    { "Data": "yourDocumentHash" }
    ```
  - Response: JSON representation of the newly mined block

- **GET /chain**
  - Response: JSON array of all blocks

## Testing

Run unit tests:
```bash
go test ./notary
```

## Contributing

1. Fork this repository  
2. Create a feature branch (`git checkout -b feature/YourFeature`)  
3. Commit your changes (`git commit -m "Add some feature"`)  
4. Push to the branch (`git push origin feature/YourFeature`)  
5. Open a Pull Request  

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Disclaimer

docNote is a proof-of-concept and not intended for production use without additional security measures (digital signatures, stronger consensus mechanisms, etc.).
