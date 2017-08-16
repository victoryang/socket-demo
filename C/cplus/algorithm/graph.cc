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

class vertex
{
public:
	vertex(int s):start(s),number(0),neighbor(NULL),next(NULL){};
	~vertex(){};
	int start;
	int number;
	Line *neighbor;
	vertex *next;
};

class Graph
{
public:
	Graph(int s, int e, int w){
		count = 1;
		head = create_new_vertex_for_graph(s,e,w);
	};

	vertex* create_new_vertex_for_graph(int s, int e, int w){
		vertex *v = new vertex(s);
		v->neighbor = new Line(e,w);
		return v;
	};
	vertex* find_vertex_in_graph(vertex* v, int s){
		if(NULL==v)
			return NULL;

		while(v){
			if(v->start == s)
				return v;
			v = v->next;
		}
		return NULL;
	};
	Line* find_line_in_graph(Line* l, int e){
		if(NULL==l)
			return NULL;
		while(l){
			if(l->end == e)
				return l;
			l = l->next;
		}
		return NULL;
	};
	bool insert_vertex_into_graph(int start, int end, int weight){
		Vertex* v=NULL;
		Line* l=NULL;
		if(NULL==head){
			head = create_new_vertex_for_graph(start, end, weight);
			head->number++;
			count++;
			return TRUE;
		};

		v = find_vertex_in_graph(head, s);
		if(v==NULL){
			v = create_new_vertex_for_graph(start, end, weight);
			v->next = head;
			head->next = v;
			head->number++;
			count++;
			return TRUE;
		};

		l = find_line_in_graph(v->neighbor, end);
		if(l) return FALSE;
		l = new Line(end, weight);
		l->next = v->neighbor;
		v->neighbor = l;
		v->number++;
		return TRUE;
	};

	bool delete_old_vertex(Vertex* v, int s){

	};

	bool delete_old_line(Line* l, int e){
	};

	bool delete_vertex_from_graph(int start, int end, int weight){

	};
	~Graph(){};
	int count;
	vertex *head;
};

int main(int argc, char const *argv[])
{
	
	return 0;
}