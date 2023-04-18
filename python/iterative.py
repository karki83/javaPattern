import socket

HOST = ''
PORT = 8888

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind((HOST, PORT))
sock.listen(5)


sock.setblocking(0)

client_connections = []

while True:
    try:
        conn, addr = sock.accept()
        conn.setblocking(0)
        client_connections.append(conn)
    except:
        pass    
    for c in client_connections:
        
        try:
            data = c.recv(1024)
            if data:
                c.sendall(data)
                print(data)
            else:
                c.close()
                client_connections.remove(c)
        except:
            pass
