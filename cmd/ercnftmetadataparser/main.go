// cmd/ercnftmetadataparser/main.go
package main

import (
"flag"
"log"
"os"

"ercnftmetadataparser/internal/ercnftmetadataparser"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ercnftmetadataparser.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
