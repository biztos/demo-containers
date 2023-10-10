# demo-containers

A collection of (arguably) useful containers for hacking on Docker,
Kubernetes, Podman, and similar container-based systems.

Nothing serious here; and no security whatsoever, use at your own risk, &c.

## Usage

Generally something like:

```bash

cd echo

make

docker run -p 8080:8080 echo

curl -i http://localhost:8080/hello/world

```

## Prerequisites

1. Go
2. (optionally) Make
3. (if using Make) upx
3. Docker (if using Make) or Podman or equivalent

