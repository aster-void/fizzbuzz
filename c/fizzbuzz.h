#include <stdlib.h>
#include <string.h>
#include "./fizzbuzz/fizz.h"
#include "./fizzbuzz/buzz.h"

char *apply(int n) {
  char *fizzstr = malloc(5 * sizeof(char));
  char *buzzstr = malloc(5 * sizeof(char));
  fizz(fizzstr, n);
  buzz(buzzstr, n);
  if (*fizzstr == '\0' && *buzzstr == '\0') {
    char *buf = malloc(10 * sizeof(char));
    sprintf(buf, "%d", n);
    // free(fizzstr);
    // free(buzzstr);
    return buf;
  } else {
    char* res = strcat(fizzstr, buzzstr);
    free(buzzstr);
    return res;
  }
}

char **range(int start, int last) {
  int len = start - last + 1;
  char **list = malloc(len * sizeof(char*));
  for (int i = 0; i < len; i++) {
    list[i] = apply(i + start);
  }
  return list;
}
