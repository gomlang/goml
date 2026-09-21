#include <errno.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdatomic.h>
#include <unistd.h>

uint64_t probe_integers(uint64_t a, uint64_t b, uint64_t c, uint64_t d,
                        uint64_t e, uint64_t f, uint64_t g, uint64_t h) {
    return a + 2*b + 3*c + 4*d + 5*e + 6*f + 7*g + 8*h;
}

double probe_mixed(int a, double b, float c, int d, double e, int f,
                   double g, int h, double i, int j, double k, int l,
                   double m, int n, double o, int p, double q, float r) {
    return a+b+c+d+e+f+g+h+i+j+k+l+m+n+o+p+q+r;
}

float probe_float(float x) { return x * -1.5f; }

static _Thread_local int thread_value;

int probe_tls(int value) {
    thread_value = value;
    usleep(1000);
    errno = value;
    for (volatile int i = 0; i < 100000; i++) {}
    return thread_value == value && errno == value;
}

int probe_stack(int depth) {
    volatile unsigned char scratch[8192];
    for (int i = 0; i < 8192; i++) scratch[i] = (unsigned char)depth;
    int value = depth ? probe_stack(depth - 1) : 0;
    return value + scratch[4096];
}

int probe_block(_Atomic uint32_t *entered, int fd) {
    atomic_store(entered, 1);
    unsigned char value;
    return read(fd, &value, 1) == 1 ? value : -1;
}

uint64_t probe_spin(uint64_t count) {
    volatile uint64_t result = 0;
    for (uint64_t i = 0; i < count; i++) result += i;
    return result;
}
