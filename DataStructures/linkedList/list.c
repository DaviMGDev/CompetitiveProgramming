#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct {
  int   value;
  Node* next;
  Node* prev;
} Node;

typedef struct {
  Node* head;
  int   size;
} List;

Node* new_node(int value) {
  Node* new  = malloc(sizeof(Node));
  new->value = value;
  return new;
}

void del_node(Node* node) {
  free(node);
  node = NULL;
}

void  append(List list, int value) {
  Node* new = new_node(value);
  if (list.head == NULL) {
    list.head = new;
    list.size = 1;
    return;
  }
  list.size++;
  Node* node = list.head;
  while (node->next) {
    node = node->next;
  }
  node->next = new;
}

void  push(List list, int value) {
  Node* new = new_node(value);
  if (list.head == NULL) {
    list.head = new;
    list.size = 1;
    return;
  }
  list.size++;
  new->next = list.head;
  list.head = new;
}

void  insert(List* list, int index, int value) {
}
Node* pop(Node* node, int index) {}
void  remove(Node* node, int index) {}
Node* search(Node* node, int value) {}
