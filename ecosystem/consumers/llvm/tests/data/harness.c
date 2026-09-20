
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
extern int64_t goml_mix(int64_t);
extern int64_t goml_absolute_mix(int64_t);
extern int64_t goml_sum(int64_t);
extern double goml_scale(double);
extern int64_t goml_struct(int64_t);
extern double goml_convert(int64_t);
int main(void) {
    for (int64_t x = -200; x <= 200; ++x) {
        int64_t n = x < 0 ? -x : x;
        printf("%" PRId64 " %" PRId64 " %" PRId64 " %.17g %" PRId64 " %.17g\n",
            goml_mix(x), goml_absolute_mix(x), goml_sum(n), goml_scale((double)x), goml_struct(x), goml_convert(x));
    }
    return 0;
}
