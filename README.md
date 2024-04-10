# envoy-extproc-crc32-check-demo-go

This repository contains a demo application written in Go that demonstrates the usage of Envoy's External Processor (ExtProc) filter to do `crc(32) check` for HTTP request.

## Overview

The Envoy ExtProc filter allows you to offload request processing logic to an external process, enabling you to customize and extend Envoy's functionality. This demo application showcases how to implement an ExtProc filter in Go.

## Features

   + Integration with Envoy's External Processor filter
   + Customizable request processing logic
   + Demonstrates handling of HTTP requests in Go
   + Simple and easy-to-understand codebase

## Getting Started

To get started with the demo application, follow these steps:

  1. Clone the repository:
     ```
     git clone https://github.com/projectsesame/envoy-extproc-crc32-check-demo-go.git
     ```

  2. Build the Go application:
     ```
     go build .
     ```

  3. Run the application:
     ```
     ./envoy-extproc-crc32-check-demo-go crc32-check --log-stream --log-phases poly "0x82f63b78"
     ```

  4. Do request:
     ```
     curl --request POST \
     --url http://127.0.0.1:8080/post \
     --data '{
      "data": "123456789",
      "crc32": "E7C41C6B",
     }'
     ```

     Field Description:

      + **data**:  the content.
      + **crc32**: crc for the data.

     PS:

      The example crc is calculated using the following settings, which are the same configuration as the standard package of go.

      + **Bit Width**:                32
      + **REFIN**:                    true
      + **REFOUT**:                   true
      + **XOROUT (HEX)**:             0xFFFFFFFF
      + **Initial Value (HEX)**:      0xFFFFFFFF
      + **Polynomial Formula (HEX)**: 0x82F63B78








## Usage

The demo application listens for incoming GRPC requests on a specified port and performs custom processing logic. You can modify the processing logic in the application code according to your requirements.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
License

This project is licensed under the Apache License Version 2.0. See the LICENSE file for details.
Acknowledgements

This demo application is based on the ExtProc filter demo(s) provided by [envoy-extproc-sdk-go](https://github.com/wrossmorrow/envoy-extproc-sdk-go). please visit it for more demos.

Special thanks to the community for their contributions and support.

## Contact

For any questions or inquiries, please feel free to reach out to us for any assistance or feedback.