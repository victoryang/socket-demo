#ifndef TREE_H
#define TREE_H
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
		Node *list;
		int length;
	} Tree;
#endif