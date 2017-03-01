#include <iostream>
#include "a.pb.h"
#include <signal.h>
#include <fstream>

int stop = 0;
sigset_t myset;
int main ()
{
	sigemptyset(&myset);
	sigaddset(&myset, SIGTERM);
	while (stop == 0)
	{
		std::cout << "in loop" << std::endl;
		sigsuspend(&myset);
	}
	return 0;
}
