import subprocess
from sys import stdout
from threading import Thread
from time import sleep
import webbrowser

# 开启consul服务


def start_consul():
    subprocess. Popen(args="consul.exe agent -dev", shell=True,
                      cwd=r"D:\C_Back\Go\Sevice\consul", stdout=stdout)

# 开启web服务


def open_consul_web():
    web_url = "http://127.0.0.1:8500"
    webbrowser.open(web_url, new=0, autoraise=True)


def grpc_user_service():
    subprocess.Popen(
        "user_service.exe", shell=True, cwd=r"D:\C_Back\Go\Shop_service\user_service", stdout=stdout)


if __name__ == '__main__':
    service_thread = Thread(target=start_consul)
    web_thread = Thread(target=open_consul_web)
    service_thread.start()
    sleep(3)
    grpc_user_service()
    sleep(2)
    web_thread.start()
