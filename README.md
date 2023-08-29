# The Go OPTional Package Helper

## Install

```bash
$ make clean install
```

## Usage

### Bash helper

It is not possible for a process to modify the parent environment, which is why
we need to introduce the `goptuse` helper function. This overwrites the `PATH`
environment variable with the output of `gopt use {parameters}`

Add to `~/.bashrc`:

```bash
# This should point to your package directory:
export GOPT_PACKAGES=$HOME/opt/packages

# This is the command (function) you'll use to update your path:
function goptuse() {
	export PATH="$(gopt use "$@")"
}
```

### List packages

```bash
# List all packages and latest version (-t for pretty tables)
$ gopt list -t
+----------------------+--------------------------------------------+----------------+
| PACKAGE              | PATH                                       | LATEST VERSION |
+----------------------+--------------------------------------------+----------------+
| aws-cli              | /home/hc/opt/packages/aws-cli              | 2.13.9         |
| aws-vault            | /home/hc/opt/packages/aws-vault            | 7.2.0          |
| cilium-cli           | /home/hc/opt/packages/cilium-cli           | 0.15.5         |
| drone-cli            | /home/hc/opt/packages/drone-cli            | 1.7.0          |
| flux2                | /home/hc/opt/packages/flux2                | 2.0.1          |
| github-cli           | /home/hc/opt/packages/github-cli           | 2.32.1         |
| go-containerregistry | /home/hc/opt/packages/go-containerregistry | 0.16.1         |
| gomplate             | /home/hc/opt/packages/gomplate             | 3.11.5         |
| helm                 | /home/hc/opt/packages/helm                 | 3.12.3         |
| k9s                  | /home/hc/opt/packages/k9s                  | 0.27.4         |
| krew                 | /home/hc/opt/packages/krew                 | 0.4.4          |
| kubectl              | /home/hc/opt/packages/kubectl              | 1.28.0         |
| kubelogin            | /home/hc/opt/packages/kubelogin            | 1.28.0         |
| kubevirt             | /home/hc/opt/packages/kubevirt             | 1.0.0          |
| logcli               | /home/hc/opt/packages/logcli               | 2.8.4          |
| pv-migrate           | /home/hc/opt/packages/pv-migrate           | 1.2.0          |
| terraform            | /home/hc/opt/packages/terraform            | 1.5.5          |
| terraform-docs       | /home/hc/opt/packages/terraform-docs       | 0.16.0         |
| terragrunt           | /home/hc/opt/packages/terragrunt           | 0.50.1         |
| velero               | /home/hc/opt/packages/velero               | 1.11.1         |
| yq                   | /home/hc/opt/packages/yq                   | 4.35.1         |
+----------------------+--------------------------------------------+----------------+

# List all terraform versions (-t for pretty tables)
$ gopt list -t terraform
+-----------+---------+---------------------------------------+
| PACKAGE   | VERSION | PATH                                  |
+-----------+---------+---------------------------------------+
| terraform | 1.0.9   | /home/hc/opt/packages/terraform/1.0.9 |
| terraform | 1.1.7   | /home/hc/opt/packages/terraform/1.1.7 |
| terraform | 1.3.7   | /home/hc/opt/packages/terraform/1.3.7 |
| terraform | 1.5.5   | /home/hc/opt/packages/terraform/1.5.5 |
+-----------+---------+---------------------------------------+
```

### Select (use) package

```bash
# If you haven't run goptuse previously, terraform won't be in the path:
$ terraform --version
bash: terraform: command not found

# No parameters means pick latest version of all found packages:
$ goptuse

# Verify that we added the correct binary to the path:
$ which terraform
/home/hc/opt/packages/terraform/1.5.5/terraform

# Now pick another one:
$ goptuse terraform:1.3.7
Terraform v1.5.5

# Verify that the new version is in the path:
$ which terraform
/home/hc/opt/packages/terraform/1.3.7/terraform

# gopt use will append to PATH by default. Add -p if you with to prepend instead
# You can also add this to the bash function above.
$ goptuse -p terraform:1.3.7
```
