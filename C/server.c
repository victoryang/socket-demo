#include <stdio.h>
#include <stdlib.h>
#include <netdb.h>
#include <errno.h>
#include <string.h>
#include <time.h>
#include <sys/epoll.h>
#include <fcntl.h>
#include <signal.h>

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
	int ret;
	// Non-blocking read
	for(;;) {
		ret = recv(fd, (void*)buff, 16, 0);
		if (ret > 0) {
			printf("%s\n", buff);
		} else if (ret == 0) {
			printf("* INFO recv finished: %s\n", strerror(errno));
			break;
		} else if (ret == -1){
			if (errno == EAGAIN || errno == EWOULDBLOCK)
				break;
		} else {
			fprintf(stderr, "Error in recv: %s\n", strerror(errno));
			break;
		}
	}

	printf("doing service ...\n");
	printf("Done\nsending back echo to client ...\n");
	send(fd, "echo", strlen("echo"), 0);
	if (epoll_ctl(epollfd, EPOLL_CTL_DEL, fd, 0) < 0)
	{
		fprintf(stderr, "DEL epoll error: %s\n", strerror(errno));
		exit(0);
	}
	printf("Sending finished.\nClose connection\n");
	close(fd);
}

void setnonblocking(int fd)
{
	int fd_set;
	fd_set = fcntl(fd, F_GETFL, 0);

	if(fd_set < 0)
	{
		fprintf(stderr, "Fd flag get fails: %s\n", strerror(errno));
		exit(0);
	}

	fd_set = fd_set | O_NONBLOCK;
	if(fcntl(fd, F_SETFL, fd_set) < 0)
	{
		fprintf(stderr, "Fd flag set fails: %s\n", strerror(errno));
		exit(0);
	}
}

int setsocketoptions(int fd)
{
	int sw = 1;
	/*
	Set SO_REUSEADDR so that we don't get error like
	address is already used
	*/
	setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &sw, sizeof(sw));
	return sw;
}

int main(int argc, char const *argv[])
{
	int ret;
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

	// set socket options
	if (ret = setsocketoptions(sockfd) < 0)
	{
		fprintf(stderr, "set socket opt error: %d\n", ret);
	}

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

	// Create an epoll instance
	epollfd = epoll_create(1);
	if( epollfd < 0)
	{
		fprintf(stderr, "epoll create fails: %s\n", strerror(errno));
		exit(-1);
	}

	// Set listening fd to epoll instance
	ev.events = EPOLLIN;
	ev.data.fd = sockfd;
	if( epoll_ctl(epollfd, EPOLL_CTL_ADD, sockfd, &ev) < 0)
	{
		fprintf(stderr, "epollctl failed: %s\n", strerror(errno));
		exit(-1);
	}

	/*Empty the events array
	  It is used when a real event comes and will be set by kernel
	  so that we can get the event by polling it
	*/
	int i;
	memset((void*)events, 0, sizeof(events));

	// main loop
	for(;;) {
		// Blocking the thread for I/O request
		ndfs = epoll_wait(epollfd, events, MAX_EVENTS, -1);
		if (ndfs < 0)
		{
			/*
			This is for strace debug, strace will send a EINTR to
			epoll_wait, which will make epoll_wait fails
			*/
			if (EINTR == errno)
				continue;
			// If not, print error and quits
			fprintf(stderr, "epoll_wait fails: %s\n", strerror(errno));
			exit(-1);
		}

		// Events happens and polling the event array
		for (i = 0; i < ndfs; i++)
		{
			if ( (events[i].events & EPOLLERR) ||
			 	 (events[i].events & EPOLLHUP) ||
			 	 (!(events[i].events & EPOLLIN)) ) {
				fprintf (stderr, "epoll error\n");
				close (events[i].data.fd);
				continue;
			}

			if ( events[i].data.fd == sockfd) {
				conn = accept(sockfd, (struct sockaddr *)&client, &len);
				if (conn < 0)
				{
					fprintf(stderr, "connection error: %s\n", strerror(errno));
					exit(-1);
				}
				/* fcntl(conn, F_SETFL, O_NONBLOCK);
				   Set the new connection as non-blocking
				   which is suggested by epoll(2) when we have set
				   EPOLLET flag in event
				   Set as non-blocking flag to avoid a blocking read/write will
				   starve a task that hanlding the multiple fd
				   EPOLLONESHOT: tell epoll to disable the associated fd that receive a
				                 bunch of data 
				*/
				setnonblocking(conn);
				/*
				Set edge trigger flag to events, combined with non-blocking model
				*/                 
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