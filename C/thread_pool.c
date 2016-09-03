#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <errno.h>
#include <assert.h>

typedef struct Worker {
	void *(*process)(void* arg);
	void *arg;
	struct Worker *next;
} CThread_worker;

typedef struct Pool {
	pthread_mutex_t queue_lock;
	pthread_cond_t queue_ready;

	CThread_worker *queue_head;
	int destory;
	pthread_t *threadid;
	int max_thread_num;
	int cur_queue_size; 
} CThread_pool;

static CThread_pool *pool = NULL;
void *my_thread_init(void *arg);

void pool_init(int max_thread_num)
{
	pool = (CThread_pool*)malloc(sizeof(CThread_pool));
	if (NULL == pool)
	{
		fprintf(stderr, "Thread pool create fails: %s\n", strerror(errno));
		exit(0);
	}

	pthread_mutex_init(&(pool->queue_lock), NULL);
	pthread_cond_init(&(pool->queue_ready), NULL);
	pool->queue_head = NULL;
	pool->max_thread_num = max_thread_num;
	pool->cur_queue_size = 0;

	pool->destory = 0;
	pool->threadid = (pthread_t*)malloc(max_thread_num * sizeof(pthread_t));

	int i = 0;
	for (i = 0; i < max_thread_num; i++)
	{
		pthread_create(&(pool->threadid[i]), NULL, my_thread_init, (void*)NULL);
	}
}

int pool_add_worker(void *(process)(void *arg), void *arg)
{
	printf("pool_add_worker NO. %d\n", *((int*)arg));
	CThread_worker * new_worker = (CThread_worker*)malloc(sizeof(CThread_worker));
	new_worker->process = process;
	new_worker->arg = arg;
	new_worker->next = NULL;

	pthread_mutex_lock(&(pool->queue_lock));
	CThread_worker *member = pool->queue_head;
	if (member) {
		while(member->next)
			member = member->next;
		member->next = new_worker;
	} else {
		pool->queue_head = new_worker;
	}

	assert(pool->queue_head != NULL);

	pool->cur_queue_size++;
	pthread_mutex_unlock(&(pool->queue_lock));
	pthread_cond_signal(&(pool->queue_ready));
	return 0;
}

int pool_destory()
{
	void *ret;
	if (pool->destory)
		return -1;
	pool->destory = 1;

	pthread_cond_broadcast(&(pool->queue_ready));

	int i;
	for (i = 0; i < pool->max_thread_num; ++i)
	{
		pthread_join(pool->threadid[i], &ret);
	}
	free(pool->threadid);
	pool->threadid = NULL;

	CThread_worker *head = NULL;
	while(pool->queue_head)
	{
		head = pool->queue_head;
		pool->queue_head = pool->queue_head->next;
		free(head);
	}
	pool->cur_queue_size = 0;
	pool->max_thread_num = 0;

	pthread_mutex_destroy(&(pool->queue_lock));
	pthread_cond_destroy(&(pool->queue_ready));

	free(pool);
	pool = NULL;
	return 0;
}

void *my_thread_init(void *arg)
{
	pthread_t tid = pthread_self();
	printf("thread 0x%x created\n", (unsigned int) tid);

	while(1) {
		pthread_mutex_lock(&(pool->queue_lock));

		while( pool->cur_queue_size == 0 && !pool->destory)
		{
			printf("pthread 0x%x is waiting ...\n", pthread_self());
			pthread_cond_wait(&(pool->queue_ready), &(pool->queue_lock));
		}

		if (pool->destory) {
			pthread_mutex_unlock (&(pool->queue_lock));
			printf ("thread 0x%x will exit\n", pthread_self ());
			pthread_exit(NULL);
		}

		printf ("thread 0x%x is starting to work/n", pthread_self ());

		assert( pool->cur_queue_size != 0);
		assert( pool->queue_head != NULL);

		pool->cur_queue_size--;
		CThread_worker *worker = pool->queue_head;
		pool->queue_head = worker->next;
		pthread_mutex_unlock(&(pool->queue_lock));

		(*(worker->process))(worker->arg);
		free(worker);
		worker = NULL;
	}

	pthread_exit((void*)NULL);
}

void* process(void* arg)
{
	printf("threadid 0x%x is processing task %d\n", pthread_self(), *(int*)arg);
	sleep(1);
	return (void*)NULL;
}

int main(int argc, char const *argv[])
{
	pool_init(3);

	int *workingnum = (int*)malloc(sizeof(int) * 10);
	int i;
	for (i = 0; i < 10; i++)
	{
		workingnum[i] = i;
		pool_add_worker(process, &workingnum[i]);
	}

	sleep(5);

	pool_destory();
	free(workingnum);
	return 0;
}