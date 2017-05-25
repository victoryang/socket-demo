#include <iostream>
#include <vector>
#include <map>
using namespace std;

vector<int> TwoSum (vector<int> input, int sum)
{
	map<int, int> m;
	vector<int> result;
	for (int i = 0; i < input.size(); ++i)
	{
		/* code */
		m[input[i]] = i;
	}

	for (int i = 0; i < input.size(); ++i)
	{
		/* code */
		int find = sum - input[i];
		if (m.find(find) != m.end()
			&& m.at(find) != i)
		{
			/* code */
			result.push_back (i+1);
			result.push_back (m[find]+1);
		}
	}
	return result;
}

int main(int argc, char const *argv[])
{
	vector<int> v;
	vector<int> r;
	v.push_back(2);
	v.push_back(7);
	v.push_back(11);
	v.push_back(15);
	r = TwoSum (v, 9);
	cout << r[0] << " "<< r[1] << endl;
	return 0;
}