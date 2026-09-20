#include <stdio.h>

typedef struct sqlite3 sqlite3;
extern int sqlite3_open(const char *, sqlite3 **);
extern int sqlite3_exec(sqlite3 *, const char *, int (*)(void *, int, char **, char **), void *, char **);
extern int sqlite3_close(sqlite3 *);
extern void sqlite3_free(void *);
extern const char *sqlite3_errmsg(sqlite3 *);

static int row(void *context, int count, char **values, char **names) {
    (void)context;
    (void)names;
    for (int index = 0; index < count; index++) {
        if (index) fputc('\t', stdout);
        fputs(values[index] ? values[index] : "NULL", stdout);
    }
    fputc('\n', stdout);
    return ferror(stdout) ? 1 : 0;
}

int main(int count, char **arguments) {
    if (count != 3) return 2;
    sqlite3 *database = NULL;
    int result = sqlite3_open(arguments[1], &database);
    if (result != 0) {
        fputs(database ? sqlite3_errmsg(database) : "SQLite open failed", stderr);
        sqlite3_close(database);
        return 1;
    }
    char *error = NULL;
    result = sqlite3_exec(database, arguments[2], row, NULL, &error);
    if (result != 0) {
        fputs(error ? error : sqlite3_errmsg(database), stderr);
        sqlite3_free(error);
    }
    int closed = sqlite3_close(database);
    return result == 0 && closed == 0 ? 0 : 1;
}
