#include <thread>
#include <iostream>

using namespace std;

void hello ()
{
	cout << "hello" << endl;
}

int main(int argc, char const *argv[])
{
	std::thread t1(hello);
	t1.join();
	cout << "test!" << endl;
	return 0;
}
