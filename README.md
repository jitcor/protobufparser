### Project Overview

`protobufparser` is a Go library designed for parsing Google Protocol Buffers (protobuf) files. It can read protobuf
files and extract data from specified fields, making it suitable for applications that need to process and analyze
protobuf data.

### Key Features

1. **Parse Protobuf Files**: Reads and parses protobuf files, extracting field data.
2. **Field Query**: Queries and returns data based on specified field positions.
3. **Error Handling**: Manages file reading and parsing errors, providing detailed information.

### Usage

1. **Installation**:
   ```bash
   go get github.com/jitcor/protobufparser
   ```

2. **Basic Example**:
   ```go
   package main

   import (
       "fmt"
       "github.com/jitcor/protobufparser"
       "io/ioutil"
   )

   func main() {
       fileData, err := ioutil.ReadFile("path/to/protobuf/file")
       if err != nil {
           panic(err)
       }

       values, err := protobufparser.PQuery(fileData, 1, 2, 4, 10, 5)
       if err != nil {
           panic(err)
       }

       for _, v := range values {
           fmt.Println(string(v.([]byte)))
       }
   }
   ```

### Contribution Guidelines

We welcome contributions, including reporting issues, suggesting features, or submitting code improvements. Please check
the project's contribution guidelines for more information.

For more details, visit the [protobufparser GitHub repository](https://github.com/jitcor/protobufparser).
