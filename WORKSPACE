# The WORKSPACE file tells Bazel that this directory is a "workspace", which is like a project root.
# The content of this file specifies all the external dependencies Bazel needs to perform a build.

####################################
# ESModule imports (and TypeScript imports) can be absolute starting with the workspace name.
# The name of the workspace should match the npm package where we publish, so that these
# imports also make sense when referencing the published package.
workspace(name = "angular_bazel_GRPC")

####################################
# Fetch external repositories containing Bazel build toolchain support.
# Bazel doesn't support transitive WORKSPACE deps, so we must install those too.

# Allows Bazel to run tooling in Node.js
http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "6139762b62b37c1fd171d7f22aa39566cb7dc2916f0f801d505a9aaf118c117f",
    strip_prefix = "rules_nodejs-0.9.1",
    url = "https://github.com/bazelbuild/rules_nodejs/archive/0.9.1.zip",
)

# The Bazel buildtools repo contains tools like the BUILD file formatter, buildifier
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "dad19224258ed67cbdbae9b7befb785c3b966e5a33b04b3ce58ddb7824b97d73",
    strip_prefix = "buildtools-b3b620e8bcff18ed3378cd3f35ebeb7016d71f71",
    # Note, this commit matches the version of buildifier in angular/ngcontainer
    url = "https://github.com/bazelbuild/buildtools/archive/b3b620e8bcff18ed3378cd3f35ebeb7016d71f71.zip",
)

# Runs the TypeScript compiler
http_archive(
    name = "build_bazel_rules_typescript",
    sha256 = "1aa75917330b820cb239b3c10a936a10f0a46fe215063d4492dd76dc6e1616f4",
    strip_prefix = "rules_typescript-0.15.0",
    url = "https://github.com/bazelbuild/rules_typescript/archive/0.15.0.zip",
)

# Used by the ts_web_test_suite rule to provision browsers
http_archive(
    name = "io_bazel_rules_webtesting",
    sha256 = "cecc12f07e95740750a40d38e8b14b76fefa1551bef9332cb432d564d693723c",
    strip_prefix = "rules_webtesting-0.2.0",
    url = "https://github.com/bazelbuild/rules_webtesting/archive/v0.2.0.zip",
)

# Runs the Sass CSS preprocessor
http_archive(
    name = "io_bazel_rules_sass",
    sha256 = "b243c4d64f054c174051785862ab079050d90b37a1cef7da93821c6981cb9ad4",
    strip_prefix = "rules_sass-0.1.0",
    url = "https://github.com/bazelbuild/rules_sass/archive/0.1.0.zip",
)

# Some of the TypeScript tooling is written in Go.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "feba3278c13cde8d67e341a837f69a029f698d7a27ddbb2a202be7a10b22142a",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.3/rules_go-0.10.3.tar.gz",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
)

# Rules for producing a GRPC gateway and translating protocol buffers to swagger definitions
http_archive(
    name = "io_grpc_grpc_java",
    sha256 = "f900380a5477bca95ecd924ab18e79e588014f9c7aba0cadc21db19d540c20af",
    strip_prefix = "grpc-java-1.13.0",
    url = "https://github.com/grpc/grpc-java/archive/v1.13.0.tar.gz",
)

# Java GRPC
load("@io_grpc_grpc_java//:repositories.bzl", "grpc_java_repositories")

grpc_java_repositories()

grpc_gateway_version = "739cd2db2d2fb68c640b39110c364a2ade7ef53b"

http_archive(
    name = "grpc_ecosystem_grpc_gateway",
    sha256 = "d3da02226e8758d72f6eef5349de741c52398a666ebfb893744f5b9a5269e67c",
    strip_prefix = "grpc-gateway-{v}".format(v = grpc_gateway_version),
    url = "https://github.com/grpc-ecosystem/grpc-gateway/archive/{v}.zip".format(v = grpc_gateway_version),
)

# Swagger Code Gen Jar for producing Angular HTTP Services
http_jar(
    name = "io_swagger_swagger_codegen_cli",
    sha256 = "4fa9c74f00fc969cc15326f95c61f6d699e434371a9d02461b4f5fdbdc7a8381",
    url = "https://oss.sonatype.org/content/repositories/snapshots/io/swagger/swagger-codegen-cli/2.4.0-SNAPSHOT/swagger-codegen-cli-2.4.0-20180611.162651-269.jar",
)

http_jar(
    name = "google_java_format",
    sha256 = "73faf7c9b95bffd72933fa24f23760a6b1d18499151cb39a81cda591ceb7a5f4",
    url = "https://github.com/google/google-java-format/releases/download/google-java-format-1.6/google-java-format-1.6-all-deps.jar",
)

####################################
# Tell Bazel about some workspaces that were installed from npm.

# The @angular repo contains rule for building Angular applications
local_repository(
    name = "angular",
    path = "node_modules/@angular/bazel",
)

# The @rxjs repo contains targets for building rxjs with bazel
local_repository(
    name = "rxjs",
    path = "node_modules/rxjs/src",
)

####################################
# Load and install our dependencies downloaded above.

load("@build_bazel_rules_nodejs//:defs.bzl", "check_bazel_version", "node_repositories", "yarn_install")

node_repositories(package_json = ["//:package.json"])

load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
load("@io_bazel_rules_webtesting//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories()

browser_repositories(
    chromium = True,
    firefox = True,
)

load("@build_bazel_rules_typescript//:defs.bzl", "ts_setup_workspace")

ts_setup_workspace()

load("@io_bazel_rules_sass//sass:sass_repositories.bzl", "sass_repositories")

sass_repositories()

#####################################
# GRPC Gateway dependencies and rules
load("@grpc_ecosystem_grpc_gateway//:repositories.bzl", "repositories")

repositories()

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()
