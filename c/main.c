#include <stdio.h>
#include "./fizzbuzz.h"

int main() {
  for (int i = 1; i <= 15; i++) {
    char *v = apply(i);
    printf("%s\n", v);
    free(v);
  }
}
