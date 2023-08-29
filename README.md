# Usage

Add to `~/.bashrc`:

```bash
export GOPT_PACKAGES=$HOME/opt/packages

function goptuse() {
	export PATH
	eval $(gopt use "$@")
}
```

List packages:

```bash
# Syntax: gopt list {package} {package..}
# Example:
gopt list kubectl
```

Add package to PATH (use):

```bash
# Syntax: goptuse {package{:version}} {package{:version}}
# Example:
goptuse kubectl:1.2.3 terraform:2.3.4 terragrunt
```
