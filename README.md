# Routriever

[![Test Status](https://github.com/jo-hoe/routriever/workflows/test/badge.svg)](https://github.com/jo-hoe/routriever/actions?workflow=test)
[![Lint Status](https://github.com/jo-hoe/routriever/workflows/lint/badge.svg)](https://github.com/jo-hoe/routriever/actions?workflow=lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/routriever)](https://goreportcard.com/report/github.com/jo-hoe/routriever)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/routriever/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/routriever?branch=main)

Collects length of given routes and provides them as business metric to prometheus.

## Planned Architecture

The deployment will run a pod, a service, and provide a service monitor to track the metric.

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
...
```

The service monitor can be consumed by Prometheus.

## How To Run

To run the service locally, you can use `docker-compose`.
Before you have to provide the api key via a file with name `secret.txt` in folder `dev`.
This file should contain the following content:

```.txt
<your_tomtom_api_key>
```

In case you do not want to have the API key stored in plain text, consider to mount the dev folder as a volume, e.g. with [VeraCrypt](https://www.veracrypt.fr/en/Home.html).

Afterwards you can run the service with

```bash
docker-compose up
```

### Optional

Run the project is using `make`. `make` is typically installed by default on Linux and Mac.

If you run on Windows, you can directly install it from [gnuwin32](https://gnuwin32.sourceforge.net/packages/make.htm) or via `winget`

```PowerShell
winget install GnuWin32.Make
```

## How to Use

You can check all `make` command by running

```bash
make help
```

### Test

You can use `make` to start the service

```bash
make test
```

## Linting

Project used `golangci-lint` for linting.

<https://golangci-lint.run/welcome/install/>

Run the linting locally by executing

```bash
golangci-lint run ./...
```
