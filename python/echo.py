import socket

HOST = ''   
PORT = 8002

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
sock.bind((HOST, PORT))
sock.listen()


while True:
    conn, addr = sock.accept()
    with conn:
        while True:
            data = conn.recv(1024)
            if not data:
                break
            print(f"Received: {data}")
            conn.sendall(data)

        
