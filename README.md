# Windows, bazel, git, plink.exe

Bazel builds can timeout when `plink.exe` is used as SSH transport agent for Git.

On some Windows environments, Users prefer to use [PuTTY](https://www.putty.org/)'s `plink.exe` as SSH connection agent. They will set `GIT_SSH=plink.exe` in their environments. Due to the non-interactive nature of the commands run from a bazel build rule, the prompt to verify an unknown host key and add it to the registry goes unnoticed, hence the build times out after some (long) time.

## Reproduce

0. Make sure the host key for `git.apache.org` is *not* in the registry under `HKEY_CURRENT_USER\SoftWare\SimonTatham\PuTTY\SshHostKeys`
1. Set `GIT_SSH=plink.exe`.
2. Run `bazel build go_default_library`.
3. Observe the build hanging at some point and eventually failing with
```
ERROR: no such package '@io_opencensus_go//plugin/ocgrpc': failed to generate BUILD files for go.opencensus.io: Timed out and referenced by '@com_github_Xjs_bazel_opencensus//:go_default_library'
ERROR: Analysis of target '//:main' failed; build aborted: Analysis failed
```

The failure is due to `plink.exe` not knowing 

## Disclaimer

I noticed this first on a repository that imported `go.opencensus.io/plugin/ocgrpc`, which leads to gazelle trying to clone `git.apache.org/thrift.git`. On our hosts, the host key for `git.apache.org` was not known to `plink.exe` (Registry: `HKEY_CURRENT_USER\SoftWare\SimonTatham\PuTTY\SshHostKeys`), leading to timeouts on Bazel builds for packages that imported `go.opencensus.io/plugin/ocgrpc` with `go_repository`.

This is a very specific edge case, but I think it might occur frequently on freshly set-up Windows hosts which choose to use `plink.exe` for Git's SSH transport.