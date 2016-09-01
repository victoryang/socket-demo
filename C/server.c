#include <stdio.h>
#include <netdb.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <sys/time.h>

#define closesocket close
int sock = 0;
struct sockaddr_in addr;

void getsocketopt(int fd, int flag)
{
	char optval[20]={0};
	socklen_t len;
	getsockopt(fd, SOL_SOCKET, flag, optval, &len);
	printf("value of %d is %s\n", flag, optval);
}

void do_printf(struct sockaddr_in addr)
{
	char buffer[16];
	inet_ntop(AF_INET, &addr.sin_addr.s_addr, buffer, sizeof(addr));
	unsigned int port = ntohs(addr.sin_port);
	printf("client address %s port %d\n", buffer, port);
}

void do_service(int fd)
{
	const char *data = "echo";
	char buff[16] = {0};
	//struct timeval lasttime;
	printf("send messages to clent %d: %s\n", fd, data);
	recv(fd, (void*)buff, 16, 0);
	int flag = atoi(buff);
	printf("%d\n", flag);
	getsocketopt(fd, flag);
	send(fd, data, strlen(data), 0);
	closesocket(fd);
}

int main(int argc, char const *argv[])
{
	printf("======= socket program started ======\n");

	// create a socket
	sock = socket(AF_INET,SOCK_STREAM,0);
	if (sock < 0)
	{
		fprintf(stderr, "socket created fails: %s\n", strerror(errno));
		exit(1);
	}

	// set address
	memset((void*)&addr, 0, sizeof(addr));
	addr.sin_family = AF_INET;
	addr.sin_port = htons(3000);
	addr.sin_addr.s_addr = INADDR_ANY;

	// bind a socket to port
	if (bind(sock, (struct sockaddr*)&addr,sizeof(addr)) < 0)
	{
		fprintf(stderr, "socket bind error: %s\n", strerror(errno));
		exit(2);
	}

	// listening
	if (listen(sock, 10) < 0)
	{
		fprintf(stderr, "socket listening error: %s\n", strerror(errno));
		exit(3);
	}

	while(1)
	{
		struct sockaddr_in clientAddr;
		socklen_t len = sizeof(clientAddr);
		int fd = accept(sock, (struct sockaddr*)&clientAddr, &len);
		if (fd < 0 )
		{
			fprintf(stderr, "accept failed: %s\n", strerror(errno));
			continue;
		}

		printf("Got connection from client fd: %d\n",fd);
		do_printf(clientAddr);
		do_service(fd);
		close(fd);	
	}

	printf("======= socket quits ======\n");
	return 0;
}