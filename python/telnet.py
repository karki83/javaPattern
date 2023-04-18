import sys
import socket
len(sys.argv)

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
sock.connect((sys.argv[1], int(sys.argv[2])))

while True:

    a = input()
    b = input()
    request = (str(a) + '\r\n' + str(b) + '\r\n\r\n').encode('ascii')
    sock.sendall(request)
    reply = sock.recv(1024)
    print(reply)
    print((reply).decode('utf-8'))