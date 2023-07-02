#include <stdio.h>
#include <stdint.h>

uint8_t SLOT[2000002];

int main() {
  int N;
  scanf("%d", &N);
  for (int i = 0; i < N; ++i) {
    int number;
    scanf("%d", &number);
    SLOT[number + 1000000] = 1;
  }

  for (int i = 2000001; i >= 0; --i) {
    if (SLOT[i]) {
      printf("%d\n", i - 1000000);
    }
  }
}
