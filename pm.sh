function pm() {
    local temp_file="/tmp/pm_last_dir"

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

    while [[ ! -f package.json ]]; do
        cd ..
        if [[ $(pwd) == "/" ]]; then
            echo "No package.json file found in the current directory or its parents."
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
