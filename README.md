# Routriever

[![Test Status](https://github.com/jo-hoe/routriever/workflows/test/badge.svg)](https://github.com/jo-hoe/routriever/actions?workflow=test)
[![Lint Status](https://github.com/jo-hoe/routriever/workflows/lint/badge.svg)](https://github.com/jo-hoe/routriever/actions?workflow=lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/routriever)](https://goreportcard.com/report/github.com/jo-hoe/routriever)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/routriever/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/routriever?branch=main)

Collects length of given routes and provides them as business metric to prometheus.
The deployment will runs a pod, a service, and provides a service monitor to track the metric.
The service monitor can be consumed by Prometheus.

## API Key

Currently the service only supports TomTom as GPS service.
You can create an API key in the [TomTom Developer Portal](https://developer.tomtom.com/).
Before you can start the service locally you need to provide the api key.
This can be done by setting is as environment variable `TOM_TOM_API_KEY`
In PowerShell you set

```PowerShell
$env:TOM_TOM_API_KEY = "<your tom tom api key>";
```

and in bash you can use

```bash
TOM_TOM_API_KEY="<your tom tom api key>"
```

Afterwards you can start the service.

## Prerequisites to run locally

Run the project is using `make`. `make` is typically installed by default on Linux and Mac.

If you run on Windows, you can directly install it from [gnuwin32](https://gnuwin32.sourceforge.net/packages/make.htm) or via `winget`

```PowerShell
winget install GnuWin32.Make
```

Futhermore you will need Docker and Python.
Python is only used to set the API key without the need to persist it.
If you do not want to use Python you may also create a file containing your API key and setting the environment variable `SECRET_FILE_PATH` to the file path of that file.

### How to Use

You can check all `make` command by running

```bash
make help
```

## How To Run Locally

To run the service locally, you can use `docker-compose`.

```bash
make start
```

### K3D

[Install K3D](https://k3d.io/#install-script) to run the service in a local kubernetes cluster.
Ensure your [turned on Kubernetes in Docker Desktop](https://docs.docker.com/desktop/kubernetes/#install-and-turn-on-kubernetes).
Run the following command to start the service in a local kubernetes cluster.

```bash
make k3d-start
```

and stop it with

```bash
make k3d-stop
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
