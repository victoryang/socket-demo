#include <stdio.h>
#include <stdlib.h>
#include <netdb.h>
#include <errno.h>
#include <string.h>
#include <time.h>
#include <sys/epoll.h>
#include <fcntl.h>

/* Initial variable */
int sockfd;
clock_t start, stop;
struct sockaddr_in sAddr;

/* epoll variable*/
#define MAX_EVENTS 10
struct epoll_event ev, events[MAX_EVENTS];
int epollfd, ndfs;
int conn;
struct sockaddr client;
socklen_t len = sizeof(client);

void do_service(int fd)
{
	char buff[16] = {0};
	printf("In do_service fd: %d\n", fd);
	recv(fd, (void*)buff, 16, 0);
	printf("%s\n", buff);
	send(fd, "echo", strlen("echo"), 0);
	if (epoll_ctl(epollfd, EPOLL_CTL_DEL, fd, 0) < 0)
	{
		fprintf(stderr, "DEL epoll error: %s\n", strerror(errno));
		exit(0);
	}
	close(fd);
}

int main(int argc, char const *argv[])
{
	start = clock();
	sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd < 0)
	{
		fprintf(stderr, "Failed to open socket: %s\n", strerror(errno));
	}

	memset((void*)&sAddr, 0, sizeof(sAddr));
	sAddr.sin_family = AF_INET;
	sAddr.sin_port = htons(3000);
	sAddr.sin_addr.s_addr = INADDR_ANY;
	// inet_pton(AF_INET, argv[1], (void*)&sAddr.sin_addr.s_addr);

	if (bind(sockfd, (struct sockaddr*)&sAddr, sizeof(sAddr)) < 0)
	{
		fprintf(stderr, "Bind error: %s\n", strerror(errno));
		exit(-1);
	}

	if (listen(sockfd, 10) < 0)
	{
		fprintf(stderr, "Listen errror%s\n", strerror(errno));
		exit(-1);
	}

	// epolling
	epollfd = epoll_create(1);
	if( epollfd < 0)
	{
		fprintf(stderr, "epoll create fails: %s\n", strerror(errno));
		exit(-1);
	}

	ev.events = EPOLLIN;
	ev.data.fd = sockfd;
	if( epoll_ctl(epollfd, EPOLL_CTL_ADD, sockfd, &ev) < 0)
	{
		fprintf(stderr, "epollctl failed: %s\n", strerror(errno));
		exit(-1);
	}

	int i;
	memset((void*)events, 0, sizeof(events));
	for(;;) {
		ndfs = epoll_wait(epollfd, events, MAX_EVENTS, -1);
		if (ndfs < 0)
		{
			fprintf(stderr, "epoll_wait fails: %s\n", strerror(errno));
			exit(-1);
		}

		for (i = 0; i < ndfs; i++)
		{
			if ( events[i].data.fd == sockfd) {
				conn = accept(sockfd, (struct sockaddr *)&client, &len);
				if (conn < 0)
				{
					fprintf(stderr, "connection error: %s\n", strerror(errno));
					exit(-1);
				}
				fcntl(conn, F_SETFL, O_NONBLOCK);
				ev.events = EPOLLIN | EPOLLET;
				ev.data.fd = conn;
				if ( epoll_ctl(epollfd, EPOLL_CTL_ADD, conn, &ev) < 0)
				{
					fprintf(stderr, "epoll error in connection: %s\n", strerror(errno));
					exit(-1);
				}
			} else {
				do_service(events[i].data.fd);
			}
		}
	}
	close(epollfd);
	return 0;
}