# pm

🗂️ Manage Commands Inside Projects.

Version : 2.0.0

## 🚀 Installation

### Via Homebrew

Wip 🚧

### Manual

You can copy the function in your shell's RC file. Alternatively, You can create a separate Bash script file and copy `pm.sh` into it. You'll need to load it at the beginning of your shell RC file (e.g., `.bashrc`, `.zshrc`, etc.).

```bash
source path/to/your/script.sh
```

Don't forget to resource your shell RC file:

```bash
source ~/.zshrc
```

## 💻 Usage

`pm` allows you to easily run and manage commands on your project, according to your <a>file reference</a> which acts like the root of your project.

```bash
# Will move you to the root of your project (file reference)
# then return you to your current working directory
pm npm install express
```

You can also return at the root of the project by simply use `pm` or the flag `--root` / `-r`.

```bash
# Will move you to the root of your project
pm --root
```

Your previous working directory is saved, so you can use `pm -` to go back to it.

```bash
# Will move you back to your previous working directory
pm -
```

Use the `--file` / `-f` flag to configure your file reference.

> By default, the value is `package.json`.

```bash
# If you're working on a Cargo-based project
pm --file cargo.lock
```

You can see your reference file by using the flag `--show` / `-s`.

```bash
# Will output your reference file
pm --show
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request with the details of your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [ ] Homebrew installation
- [ ] apt installation
- [ ] Add options to navigate even faster in your project (wip...)
- [ ] Upload ASCII art
- [x] Rewrite script in `Go`

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
