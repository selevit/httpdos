CC=gccgo
CFLAGS=-DNDEBUG -O2 -pipe

default:
	$(CC) -Wall httpdos.go -o httpdos $(CFLAGS) 
clean:
	rm -r httpdos
