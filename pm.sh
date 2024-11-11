function pm() {
    local current_dir=$(pwd)

    while [[ ! -f package.json ]]; do
        cd ..
        if [[ $(pwd) == "/" ]]; then
            echo "Any package.json file found in the current directory or its parents"
            cd "$current_dir"
            return 1
        fi
    done

    if [[ $# -eq 0 ]]; then
        echo "Project's root: $(pwd)"
    else
        echo "Command executed in $(pwd)"
        "$@"

        cd "$current_dir"
    fi
}