#include <iostream>
//#include "tree.h"
using namespace std;
#define MAX 9
int array[MAX] = {3,4,7,2,9,1,6,5,8};

template<class T>
class node
{
public:
	node():data(0),lchild(NULL),rchild(NULL){};
	node(const T &value, node<T> *lch = NULL, node<T> *rch = NULL):data(value),lchild(lch),rchild(rch){
		cout << "insert value: " << value << endl;
	};
	~node(){};
	T data;
	node<T> *lchild;
	node<T> *rchild;
};

template<class T>
class binary_tree
{
public:
	void insertdata(node<T> *t, T v){
		if(t == NULL){
			t = new node<T>(v,NULL,NULL);
		} else if(t->data > v){
			insertdata(t->lchild, v);
		} else {
			insertdata(t->rchild, v);
		};
	};
	binary_tree(T *a, int len):length(len){
		for (int i = 0; i < len; ++i)
		{
			insertdata(tail, a[i]);
		}
	};
	~binary_tree(){};
private:
	node<T> *tail = NULL;
	int length;
};

int main(int argc, char const *argv[])
{
	binary_tree<int> BT(array, MAX);
	return 0;
}