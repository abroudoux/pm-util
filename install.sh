#!/bin/bash

detect_shell_rc_file() {
    case "$SHELL" in
        */zsh)
            echo "$HOME/.zshrc"
            ;;
        */bash)
            if [[ -f "$HOME/.bashrc" ]]; then
                echo "$HOME/.bashrc"
            else
                echo "$HOME/.bash_profile"
            fi
            ;;
        */fish)
            echo "$HOME/.config/fish/config.fish"
            ;;
        *)
            echo ""
            ;;
    esac
}

add_to_rc_file() {
    local rc_file="$1"
    local pm_path="$2"

    if grep -q "export PATH=.*$pm_path" "$rc_file" 2>/dev/null; then
        echo "Path already added to $rc_file"
    else
        echo "export PATH=\"$pm_path:\$PATH\"" >> "$rc_file"
        echo "Path added to $rc_file"
    fi
}

install_pm() {
    local install_dir="/usr/local/bin"
    local rc_file
    local pm_path

    sudo cp pm "$install_dir"
    sudo chmod +x "$install_dir/pm"

    rc_file=$(detect_shell_rc_file)

    if [[ -z "$rc_file" ]]; then
        echo "Unsupported shell. Add $install_dir to your PATH manually."
        exit 1
    fi

    pm_path="$install_dir"
    add_to_rc_file "$rc_file" "$pm_path"

    echo "Reloading shell configuration..."
    source "$rc_file"
    echo "Installation complete. You can now run 'pm' from anywhere."
}

install_pm
