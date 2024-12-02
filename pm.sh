function pm() {
    local temp_file="/tmp/pm_last_dir"
    local config_file="/tmp/pm_config_file"
    local target_file="package.json"

    if [[ -f "$config_file" ]]; then
        target_file=$(cat "$config_file")
    fi

    if [[ "$1" == "--file"]]; then
        echo "Current target file: $target_file"
        return
    fi

    if [[ "$1" == "--file" && -n "$2" ]] || [[ "$1" == "-f" && -n "$2" ]]; then
        target_file="$2"
        echo "$target_file" > "$config_file"
        echo "Configuration updated: target file is now '$target_file'"
        return
    fi

    if [[ "$1" == "-" ]]; then
        if [[ -f "$temp_file" ]]; then
            local last_dir=$(cat "$temp_file")
            cd "$last_dir" || return 1
            echo "Returned to last directory: $last_dir"
        else
            echo "No previous directory saved."
        fi
        return
    fi

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
