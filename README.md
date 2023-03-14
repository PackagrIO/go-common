
```bash
docker pull ghcr.io/packagrio/packagr-dev:master
docker run --rm -it \
  -v $(pwd):/go/src/github.com/packagrio/go-common \
  -w /go/src/github.com/packagrio/go-common \
  ghcr.io/packagrio/packagr-dev:master bash


```
