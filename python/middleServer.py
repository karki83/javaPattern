import socket

HOST = ''
PORT = 8888

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind((HOST, PORT))
sock.listen(5)
client_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client_sock.connect(('192.168.33.10', 8000))
sock.setblocking(0)

while True:
    try:
        client_conn, client_addr = sock.accept()
        data = client_conn.recv(1024)
        if data:
            print("middler server recieved data")
            client_sock.sendall(data)
            recv = client_sock.recv(1024)
            if recv:
                print("send data recieved from server to client")
                client_conn.sendall(recv)
        client_conn.close()
    except:
        pass
