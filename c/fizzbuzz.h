#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "./fizzbuzz/fizz.h"
#include "./fizzbuzz/buzz.h"

const int MAX_STR_LEN = 10;
const int MAX_STR_SIZE = MAX_STR_LEN * sizeof(char);

char *apply(int n) {
  char *fizzstr = malloc(MAX_STR_SIZE);
  char *buzzstr = malloc(MAX_STR_SIZE);
  fizz(fizzstr, n);
  buzz(buzzstr, n);
  if (*fizzstr == '\0' && *buzzstr == '\0') {
    char *buf = malloc(MAX_STR_SIZE);
    sprintf(buf, "%d", n);
    free(fizzstr);
    free(buzzstr);
    return buf;
  } else {
    char* res = strcat(fizzstr, buzzstr);
    free(buzzstr);
    return res;
  }
}

char **range(int start, int last) {
  int len = last - start + 1;
  char **list = malloc(len * sizeof(void*));
  for (int i = 0; i < len; i++) {
    list[i] = apply(i + start);
  }
  return list;
}
