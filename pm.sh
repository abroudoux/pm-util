function pm() {
    local temp_file="/tmp/pm_last_dir"
    local config_file="/tmp/pm_config_file"
    local target_file="package.json"

    [[ -f "$config_file" ]] && target_file=$(cat "$config_file")

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --file|-f)
                if [[ -n "$2" ]]; then
                    target_file="$2"
                    echo "$target_file" > "$config_file"
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
            -)
                if [[ -f "$temp_file" ]]; then
                    local last_dir=$(cat "$temp_file")
                    cd "$last_dir" || return 1
                    echo "Returned to last directory: $last_dir"
                else
                    echo "No previous directory saved."
                fi
                return
                ;;
            --help|-h)
                echo "Usage: pm [options] [command]"
                echo "Options:"
                echo "  --file, -f <file>     Set or display the target file (default: package.json)"
                echo "  --show, -s            Show the current target file"
                echo "  -                     Return to the last saved directory"
                echo "  --help, -h            Show this help message"
                echo "  <command>             Execute a command in the project's root directory"
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

    echo "$current_dir" > "$temp_file"

    if [[ $# -eq 0 ]]; then
        echo "Project's root: $(pwd)"
    else
        echo "Command executed in $(pwd)"
        "$@"

        cd "$current_dir"
    fi
}