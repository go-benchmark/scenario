# FS Benchmark 

FS Benchmark  is simulation of users of FS Benchmark

## Installation

Use the go module to install FS Benchmark .

```go
go get github.com/go-benchmark/scenario
```

## Usage

```go
import (
    "github.com/go-benchmark/scenario/user"
    "github.com/go-benchmark/scenario/device"
    "github.com/go-benchmark/scenario/config"
    "github.com/go-benchmark/scenario/service"
)

u := service.NewUser(opts)

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)