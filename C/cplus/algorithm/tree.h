#ifndef TREE_H
#define TREE_H
typedef struct node {
	int value;
	struct node *left;
	struct node *right;
} Node;
#endif