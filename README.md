# ErcNftMetadataParser

## Description



## Features

- Parses ERC-721 and ERC-1155 metadata standards, including JSON and image formats.
- Validates metadata schemas against pre-defined JSON schema definitions to ensure data integrity.
- Caches parsed metadata locally with configurable expiration policies to minimize external API calls.
- Retrieves metadata from decentralized storage solutions like IPFS and Arweave using content addressing.
- Normalizes metadata attributes into a consistent data structure for simplified querying and analysis.
- Supports custom metadata attribute extraction through configurable XPath or JSONPath expressions.
- Detects and mitigates common metadata poisoning attacks by verifying digital signatures when available.
- Integrates with popular NFT marketplaces APIs (e.g., OpenSea, Rarible) to fetch missing metadata.
## Installation

```bash
go get github.com/uhsr/ErcNftMetadataParser
```

## Usage

```go
package main

import (
    "ercnftmetadataparser/internal/ercnftmetadataparser"
)

func main() {
    app := ercnftmetadataparser.NewApp(false)
    app.Run()
}
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
