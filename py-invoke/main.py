import ctypes
import json as j
from sanic import Sanic, Request, json


# класс обработки ошибки из ответа (go_app.so)
class Error(ctypes.Structure):
    _fields_ = [('err', ctypes.c_char_p)]

    def __del__(self):
        # Можем вызвать del_error с ошибкой None,
        # таким образом можно избежать overhead на вызов, когда в этом нет необходимости.
        if self.err is not None:
            del_error(self)

    def raise_if_err(self):
        if self.err is not None:
            raise Exception(self.err.decode())


# класс для парсинга ответа из go_app.so
class InvokeResult(ctypes.Structure):
    _fields_ = [
        ('result', ctypes.c_char_p),
        ('err', Error),
    ]


# читаем файл
lib = ctypes.CDLL('./go_app.so')

# функция invoke, импортируемая из go логики
invoke = lib.invoke
# С типы
invoke.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_longlong]
invoke.restype = InvokeResult

# функция del_error, импортируемая из go логики
del_error = lib.delError
del_error.argtypes = [Error]

# создание буфера
buf_size: int = 20 << 20  # 20mb
buf: ctypes.Array[ctypes.c_char] = ctypes.create_string_buffer(buf_size)

# with open(os.path.join("input.json"), "r") as f:
#     payload = json.load(f)
#
# resp: InvokeResult = invoke(json.dumps(payload).encode('utf-8'), buf, buf_size)
# resp.err.raise_if_err()
# print(resp.result.decode())

# сервис
app = Sanic("Go_from_Py")


# хендлер
@app.post("/execute")
async def handler(request: Request):
    resp: InvokeResult = invoke(j.dumps(request.json).encode('utf-8'), buf, buf_size)
    resp.err.raise_if_err()

    return json(j.loads(resp.result.decode()))
