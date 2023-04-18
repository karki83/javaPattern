import socket
import os

HOST = ''
PORT = 8006

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind((HOST, PORT))
sock.listen(5)


while True:
    try:
        conn, addr = sock.accept()
        pid = os.fork()
        if pid == 0:
            sock.close()
            while True:
                try:
                    data = conn.recv(1024)
                    if data:
                        conn.sendall(data)
                        print(data)
                    else:
                        conn.close()
                        break
                except:
                    pass
        else:  
            conn.close() 
    except:
        pass
