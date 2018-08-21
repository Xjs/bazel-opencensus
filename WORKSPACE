workspace(name = "com_github_Xjs_bazel_opencensus")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.0/rules_go-0.15.0.tar.gz"],
    sha256 = "56d946edecb9879aed8dff411eb7a901f687e242da4fa95c81ca08938dd23bb4",
)
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

load("@bazel_gazelle//:def.bzl", "go_repository")

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    urls = ["https://github.com/census-instrumentation/opencensus-go/archive/v0.14.0.zip"],
    strip_prefix = "opencensus-go-0.14.0",
    sha256 = "6e79f8073b26fb99cb71849b58749c2ffc591aa975edc09f4a81578c219c20a4",
)

