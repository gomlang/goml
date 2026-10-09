#ifndef GOML_C_SAMPLE_H
#define GOML_C_SAMPLE_H
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#define SAMPLE_MASK UINT64_C(18446744073709551615)

typedef enum { SampleZero = 0, SampleOne = 1 } SampleMode;
static int64_t sample_signed(int64_t value) { return value; }
static uint64_t sample_unsigned(uint64_t value) { return value; }
static double sample_float(double value) { return value; }
static _Bool sample_bool(_Bool value) { return value; }
static SampleMode sample_mode(SampleMode value) { return value; }

typedef struct SampleCounterData { int64_t value; } *SampleCounter;
static int sample_live_strings;

static SampleCounter sample_create(int64_t value) {
    SampleCounter counter = malloc(sizeof(*counter));
    if (counter != NULL) counter->value = value;
    return counter;
}

static void sample_destroy(SampleCounter counter) { free(counter); }
static int64_t sample_read(SampleCounter counter) { return counter->value; }
static int sample_open(SampleCounter *counter) {
    *counter = sample_create(42);
    return *counter == NULL ? 1 : 0;
}

static void sample_divide(int32_t number, int32_t *quotient, int32_t *remainder) {
    *quotient = number / 3;
    *remainder = number % 3;
}

static void sample_invert(unsigned char *data, size_t size) {
    for (size_t index = 0; index < size; index++) data[index] ^= 255;
}

static uint64_t sample_sum(const unsigned char *data, uint8_t size) {
    uint64_t result = 0;
    for (size_t index = 0; index < size; index++) result += data[index];
    return result;
}

static char *sample_greet(const char *name) {
    size_t length = strlen(name);
    char *result = malloc(length + 1);
    if (result != NULL) {
        memcpy(result, name, length + 1);
        sample_live_strings++;
    }
    return result;
}

static void sample_free(void *text) {
    if (text != NULL) {
        sample_live_strings--;
        free(text);
    }
}

static const char *sample_null_text(void) { return NULL; }
static const char *sample_raw_text(void) { return "\xff"; }
static char *sample_long_text(void) { return sample_greet("too long"); }
static int sample_out_text(char **text) { *text = sample_greet("output"); return 7; }
static char *sample_two_texts(char **text) { *text = sample_greet("output"); return sample_greet("first"); }
static int sample_allocations(void) { return sample_live_strings; }
#endif
