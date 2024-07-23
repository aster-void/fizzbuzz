#include <string.h>

int is_buzz(int n) {
  return n % 5 == 0;
}

void buzz(char* buf, int n) {
  if (is_buzz(n)) {
    memcpy(buf, "Buzz", strlen("Buzz") + 1);
  } else {
    memcpy(buf, "", strlen("") + 1);
  }
}
