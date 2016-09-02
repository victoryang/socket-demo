#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <netdb.h>
#include <errno.h>

int sockfd;
struct sockaddr_in addr;

int main(int argc, char const *argv[])
{
	sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd < 0)
	{
		fprintf(stderr, "socket create fails: %s\n", strerror(errno));
	}

	memset((void *)&addr, 0, sizeof(addr));
	addr.sin_family = AF_INET;
	addr.sin_port = htons(atoi(argv[2]));
	inet_pton(AF_INET, argv[1], (void*)&addr.sin_addr.s_addr);
	// addr.sin_addr.s_addr = INADDR_ANY;

	socklen_t len = sizeof(addr);
	if (connect(sockfd, (struct sockaddr*)&addr, len) < 0)
	{
		fprintf(stderr, "Error while connecting to server: %s\n",strerror(errno) );
	}
	printf("%s\n", argv[3]);
	send(sockfd, argv[3], strlen(argv[3]), 0);

	char buff[16];
	recv(sockfd, (void*)buff, 16, 0);
	printf("receive from server %s\n", buff);

	return 0;
}