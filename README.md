## Project description

Example of how to call **_GO_** code from **_Python_**.

## Environment

- Golang 1.19.2
- Python 3.9.13

## Golang code prepare

Folder `go-app`.

Golang script was compiled into an executable file `.so` using the command:

```bash
go build -buildmode=c-shared -o go_app.so main.go
```

## Python invoker

Folder `py-invoke`.

To call the GO code from Python was used Python library `ctypes` - https://pypi.org/project/ctypes/

To create the invoke service was used Router `Sanic` - https://sanic.dev/en/

To start the service run the command in the terminal:

```bash
sanic main:app --host=0.0.0.0 --port=8000 --workers=1
```

To get the calculation send a POST request to http://localhost:8000/execute. 

The example of request JSON you con find in directory:
```bash
py-invoke/input.json
```