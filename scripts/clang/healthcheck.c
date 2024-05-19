#include <stdio.h>
#include <stdlib.h>

#define MAX_COMMAND_LENGTH 256

int main(void) {
    const char *process_name = "kubernetes-playground";

    // Construct the command to check if the process is running
    char command[MAX_COMMAND_LENGTH];
    snprintf(command, sizeof(command), "ps aux | grep -v grep | grep -q %s", process_name);

    // Run the command
    int result = system(command);
    if (result == 0) {
        printf("Success: Process %s is running\n", process_name);
        return 0;
    } else {
        printf("Error: Process %s is not running\n", process_name);
        return 1;
    }
}
