#include <iostream>
//#include "tree.h"
using namespace std;
#define MAX 9
int array[MAX] = {3,4,7,2,9,1,6,5,8};

class node
{
public:
	node():data(0),lchild(NULL),rchild(NULL){};
	node(const int &value, node *lch = NULL, node *rch = NULL):data(value),lchild(lch),rchild(rch){
		cout << "insert value: " << value << endl;
	};
	~node(){};
	int data;
	node *lchild;
	node *rchild;
};


class binary_tree
{
public:
	void insertdata(node *t, int v){
		cout << "in insert data" << endl;
		if(t == NULL){
			t = new node(v,NULL,NULL);
		} else if(t->data > v){
			insertdata(t->lchild, v);
		} else {
			insertdata(t->rchild, v);
		};
	};
	binary_tree(int *a, int len):length(len){
		for (int i = 0; i < len; ++i)
		{
			cout << "i: " <<i<< endl;
			insertdata(tail, a[i]);
		}
	};
	~binary_tree(){};
private:
	node *tail = NULL;
	int length;
};

int main(int argc, char const *argv[])
{
	binary_tree BT(array, MAX);
	return 0;
}