# apple-store-exporter

## Usage
```bash
go build cmd/expoter
./export
```

## Configuration
### flags
1. --config
    path of config file, default is config.yaml
2. --bind
    address to bind prometheus exporter, default is 2122

### config file
```yaml
endpoint: https://www.apple.com.cn
timeout: 3s

interval: 15s
shops:
  - R532 #杭州万象城
  - R471 #杭州西湖
  - R688 #苏州

models:
  - MLT63CH/A # white 128G
  - MLTC3CH/A # white 256G
  - MLTG3CH/A # white 512G
  - MLTL3CH/A # white 1024G
  - MLT83CH/A # blue 128G
  - MLTE3CH/A # blue 256G
  - MLTJ3CH/A # blue 512G
  - MLTN3CH/A # blue 1024G
  - MLT53CH/A # black 128G
  - MLT93CH/A # black 256G
  - MLTF3CH/A # black 512G
  - MLTK3CH/A # black 1024G
```