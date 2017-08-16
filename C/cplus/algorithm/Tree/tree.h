#ifndef TREE_H
#define TREE_H
	#define LINK_VERSION
	#define MAX 9
	int array[MAX] = {3,4,7,2,9,1,6,5,8};
	#ifdef LINK_VERSION
		typedef struct node {
			int value;
			struct node *left;
			struct node *right;
		} Node;
		
	#endif
	#ifdef ARRAY_VARSION
		typedef struct node {
			int value;
			int left;
			int right;
		} Node;
	#endif
	
	typedef struct tree {
		Node *root;
		int len;
	} Tree;
	
	#ifdef LINK_VERSION
		// for Tree creation
	#endif
#endif