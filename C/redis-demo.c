#include <stdio.h>
#include <hiredis/hiredis.h>
#include <hiredis/async.h>
#include <errno.h>
#include <stdlib.h>
#include <time.h>

/*
	Compile it with command:
	gcc redis.c -l hiredis -o redis
*/

int main(int argc, char const *argv[])
{
	const char *ip = (argc > 1) ? argv[1] : "127.0.0.1";
	int port = (argc > 2) ? atoi(argv[2]) : 6379;
	struct timeval tv = {1, 500000};
	redisReply *reply;
	redisContext *rc = redisConnectWithTimeout(ip, port, tv);
	if ( rc == NULL || rc->err ) {
		if (rc) {
			printf("Connection error: %s\n", rc->err);
			redisFree(rc);
		} else {
			printf("Can not connect to redis\n");
		}
		exit(1);
	}

	printf("Connect to server successfully\n");
	reply = redisCommand(rc, "PING");
	printf("PING: %s\n", reply->str);
	freeReplyObject(reply);
	redisFree(rc);
	return 0;
}