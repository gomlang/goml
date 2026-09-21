#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>

extern int64_t nested(int64_t);
extern int64_t edited(int64_t);

int main(void) {
    for (int64_t n = -16; n <= 48; ++n) {
        printf("%" PRId64 " %" PRId64 "\n", nested(n), edited(n));
    }
    return 0;
}
