#include <iostream>

using namespace std;

class Line
{
public:
	Line(int e,int w):end(e),weight(w),next(NULL){};
	~Line(){};
	int end;
	int weight;
	Line *next;
};

class Vectex
{
public:
	Vectex(int s):start(s),number(0),neighbor(NULL),next(NULL){};
	~Vectex(){};
	int start;
	int number;
	Line *neighbor;
	Vectex *next;
};

class Graph
{
public:
	Graph(int s, int e, int w){
		count = 1;
		head = create_new_vectex_for_graph(s,e,w);
	};

	Vectex* create_new_vectex_for_graph(int s, int e, int w){
		Vectex *v = new Vectex(s);
		v->neighbor = new Line(e,w);
		return v;
	};
	Vectex* find_vectex_in_graph(Vectex* v, int s){
		if(NULL==v)
			return NULL;

		while(v){
			if(v->start == s)
				return v;
			v = v->next;
		}
		return NULL;
	}
	~Graph(){};
	int count;
	Vectex *head;
};

int main(int argc, char const *argv[])
{
	
	return 0;
}