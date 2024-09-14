#!/bin/bash

output_file="all_files_content.txt"

> "$output_file"

function append_file_contents() {
    local dir="$1"

    for entry in "$dir"/*; do
        if [ -d "$entry" ]; then
            append_file_contents "$entry"
        elif [ -f "$entry" ]; then
            echo "// ${entry#$PWD/}" >> "$output_file"
            cat "$entry" >> "$output_file"
            echo -e "\n\n" >> "$output_file"
        fi
    done
}

append_file_contents "."

echo "All file contents have been written to $output_file."
