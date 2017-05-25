#include <iostream>
#include <map>
#include <string>

using namespace std;

int main(int argc, char const *argv[])
{
	map<string, int> dict;
	string s;
	while(cin >> s){
		++dict[s];
	}
	map<string, int>::iterator it = dict.begin();
	for (; it != dict.end(); ++it)
	{
		/* code */
		cout << it->first << ": " << it->second << endl;
	}
	return 0;
}