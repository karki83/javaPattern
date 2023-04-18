import socket
import os

HOST = ''
PORT = 8000

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
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
                    request = data.decode('ascii')
                    print('req', request)
                    if not request:
                        break
                    path = request.split(' ')[1]
                    print('reqSplit', path)
                    if path == '/' or path == '/index.html':
                        filename = 'index.html'
                    else:
                        filename = path[1:]
                    if os.path.exists(filename):
                        with open(filename, 'r') as f:
                            response = f.read()
                        header = 'HTTP/1.1 200 OK\r\nContent-Length: ' + str(len(response)) + '\r\n\r\n'
                        conn.sendall((header + response).encode('ascii'))
                    else:
                        response = 'File not found'
                        header = 'HTTP/1.1 404 Not Found\r\nContent-Length: ' + str(len(response)) + '\r\n\r\n'
                        conn.sendall((header + response).encode('ascii'))
                except:
                    break
            conn.close()
            os._exit(0)
        else:
            conn.close()
    except:
        pass
