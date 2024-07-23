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

## Secrets

For local development it is excepted to have a `.env` file in folder `dev`.
You can use applications such a [VeraCrypt](https://www.veracrypt.fr/en/Home.html) to encrypt the file.

The .env file is excepted to have the following content:

```.env
TOMTOM_API_KEY=<your key>
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
