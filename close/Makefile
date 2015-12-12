all:handshake

handshake:
	go build -o c close.go
	./c server&
	./c client
	#sudo tcpdump -i lo0 -nnN port 10234

clean:
	@killall c
	rm -f c
