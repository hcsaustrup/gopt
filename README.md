# Usage

## Bash helper

Add to `~/.bashrc`:

```bash
export GOPT_PACKAGES=$HOME/opt/packages

function goptuse() {
	export PATH
	eval $(gopt use "$@")
}
```

## List packages

```bash
# Syntax: gopt list {-t} {package} {package..}
# Example:
gopt list kubectl
```

Add `-t` for pretty table formatting.

## Add package to PATH (use)

```bash
# Syntax: goptuse {-p} {package{:version}} {package{:version}}
# Example:
goptuse kubectl:1.2.3 terraform:2.3.4 terragrunt
```

By default, package paths will be *appended* to `PATH`. Add `-p` to *prepend* instead.
