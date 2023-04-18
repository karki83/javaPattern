import socket
import os

HOST = ''
PORT = 8000

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind((HOST, PORT))
sock.listen(5)
sock.setblocking(0)


while True:
    try:
        conn, addr = sock.accept()
        pid = os.fork()
        if pid == 0:
            sock.close()
            while True:
                try:
                    data = conn.recv(1024)
                    test = open('index.html', 'r')
                    # d = f.read()
                    print('read', data)
                    if data:
                        size = os.stat('index.html')
                        header = 'HTTP/1.1 200 OK\r\nContent-Length: 1000\r\n\r\n'.encode('ascii')
                        response = ''
                        
                        for i in range(0, size.st_size, 64):
                            response += test.read(i)
                            print('resp', response)
                        conn.sendall((response).encode('ascii'))
                        conn.sendall((response).encode('ascii'))
                        
                        # conn.send(response.encode())
                        test.close()
                    else: 
                        conn.close()
                        break
                except:
                    pass
        else:
            conn.close() 
    except:
        pass
