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

	memset((void *)&addr, sizeof(addr), 0);
	addr.sin_family = AF_INET;
	addr.

	return 0;
}