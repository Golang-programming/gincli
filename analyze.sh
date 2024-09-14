#!/bin/bash

# Define the output file
output_file="all_files_content.txt"

# Start fresh
> "$output_file"

# Function to recursively go through directories and append file contents
function append_file_contents() {
    local dir="$1"

    # Loop through the files and directories in the given directory
    for entry in "$dir"/*; do
        if [ -d "$entry" ]; then
            # If it's a directory, call this function recursively
            append_file_contents "$entry"
        elif [ -f "$entry" ]; then
            # If it's a file, append its path and contents to the output file
            echo "// ${entry#$PWD/}" >> "$output_file"
            echo "<content of $(basename "$entry")>" >> "$output_file"
            cat "$entry" >> "$output_file"
            echo -e "\n\n" >> "$output_file"
        fi
    done
}

# Start the process from the current directory
append_file_contents "."

echo "All file contents have been written to $output_file."
