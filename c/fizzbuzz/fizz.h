#include <string.h>

int is_fizz(int n) {
  return n % 3 == 0;
}

void fizz(char *buf, int n) {
  if (is_fizz(n)) {
    memcpy(buf, "Fizz", strlen("Fizz") + 1);
  } else {
    memcpy(buf, "", strlen("") + 1);
  }
}
