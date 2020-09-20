# FS Benchmark Service 

FS Benchmark Service is simulation of services of FS Benchmark

## Installation

Use the go module to install FS Benchmark Service.

```go
go get github.com/go-benchmark/scenario/service
```

## Usage

```go
import (
    "github.com/go-benchmark/scenario/service"
)

realtimeLength := 1
_s := service.NewService(s,realtimeLength)

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)