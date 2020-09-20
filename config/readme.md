# FS Benchmark Configuration 

FS Benchmark configuration is store all testcase of config for fs benchmark

## Installation

Use the go module to install fs config.

```go
go get github.com/go-benchmark/scenario/config
```

## Usage

```go
import (
    "github.com/go-benchmark/scenario/config"
)
vu := 1
options,err := config.ConfigureOptions(vu)

if err != nil {
    // handle error here
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)