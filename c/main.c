#include <stdio.h>
#include <stdlib.h>
#include "./fizzbuzz.h"

int main() {
  char **list = range(1, 15);
  for (int i = 0; i < 15; i++) {
    printf("%s\n", list[i]);
  }
  free(list);
}
