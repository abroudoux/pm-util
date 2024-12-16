#!/bin/bash

local last_working_dir="/tmp/pm_last_working_dir"
local reference_file="/tmp/pm_reference_file"

if [[ $target_file == "" ]]; then
    echo "Error: target file is not set. Use --file to set it." >&2
    return 1
fi

while [[ $# -gt 0 ]]; do
    case "$1" in
        --root|-r)
            while [[ ! -f "$target_file" ]]; do
                cd ..
                if [[ $(pwd) == "/" ]]; then
                    echo "No $target_file file found in the current directory or its parents."
                    cd "$current_dir"
                    return 1
                fi
            done
            return
            ;;
        -)
            if [[ -f "$last_working_dir" ]]; then
                local last_dir=$(cat "$last_working_dir")
                cd "$last_dir" || return 1
                echo "Returned to last directory: $last_dir"
            else
                echo "No previous directory saved."
            fi
            return
            ;;
        --file|-f)
            if [[ -n "$2" ]]; then
                target_file="$2"
                echo "$target_file" > "$reference_file"
                echo "Configuration updated: target file is now '$target_file'"
                return
            else
                echo "Error: --file requires a value." >&2
                return 1
            fi
            ;;
        --show|-s)
            echo "Current target file: $target_file"
            return
            ;;
        --version|-v)
            cat ./ascii.txt
            echo ""
            version=$(jq -r '.version' ./package.json) 
            echo "$version"
            return
            ;;
        --help|-h)
            echo "Usage: pm [options] [command]"
            echo "Options:"
            echo "  --root, -r            Change to the project's root directory"
            echo "  <command>             Execute a command in the project's root directory then return to the current directory"
            echo "  -                     Return to the last saved directory"
            echo "  --file, -f <file>     Set or display the target file (default: package.json)"
            echo "  --show, -s            Show the current target file"
            echo "  --help, -h            Show this help message"
            return
            ;;
        *)
            echo "Unknown option: $1" >&2
            echo "Use --help or -h for usage information."
            return 1
            ;;
    esac
    shift
done

local current_dir=$(pwd)

while [[ ! -f "$target_file" ]]; do
    cd ..
    if [[ $(pwd) == "/" ]]; then
        echo "No $target_file file found in the current directory or its parents."
        cd "$current_dir"
        return 1
    fi
done

echo "$current_dir" > "$last_working_dir"

if [[ $# -eq 0 ]]; then
    echo "Project's root: $(pwd)"
else
    echo "Command executed in $(pwd)"
    "$@"

    cd "$current_dir"
fi