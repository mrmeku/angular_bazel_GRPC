# The WORKSPACE file tells Bazel that this directory is a "workspace", which is like a project root.
# The content of this file specifies all the external dependencies Bazel needs to perform a build.

####################################
# ESModule imports (and TypeScript imports) can be absolute starting with the workspace name.
# The name of the workspace should match the npm package where we publish, so that these
# imports also make sense when referencing the published package.
workspace(name = "angular_bazel_example")

# This rule is built-into Bazel but we need to load it first to download more rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_jar", "http_archive")

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "895c2b0d1480834808216fb36cd4bf21517954cb966c893ee42421dfefcfd4bc",
    strip_prefix = "rules_nodejs-1f6d878a9ea3a095291f66e3d1a0f6b4641f5159",
    url = "https://github.com/bazelbuild/rules_nodejs/archive/1f6d878a9ea3a095291f66e3d1a0f6b4641f5159.zip",
)

http_archive(
    name = "build_bazel_rules_typescript",
    sha256 = "5a31a67550d2a92270599e90a221afe229f6e1863bf2eff22e9750c6ecf45978",
    strip_prefix = "rules_typescript-2963b55370b21d545d0ac0f30fca9dc74a0f5538",
    url = "https://github.com/bazelbuild/rules_typescript/archive/2963b55370b21d545d0ac0f30fca9dc74a0f5538.zip",
)

# The @angular repo contains rule for building Angular applications
http_archive(
    name = "angular",
    sha256 = "c480904802d62ce63a4955fd8679371a0d9620131c1c424c8786429f4e8ac77e",
    strip_prefix = "angular-7.1.3",
    url = "https://github.com/angular/angular/archive/7.1.3.zip",
)

# The @rxjs repo contains targets for building rxjs with bazel
http_archive(
    name = "rxjs",
    sha256 = "72b0b4e517f43358f554c125e40e39f67688cd2738a8998b4a266981ed32f403",
    strip_prefix = "package/src",
    url = "https://registry.yarnpkg.com/rxjs/-/rxjs-6.3.3.tgz",
)

# Angular material
# NOTE: using a `7.1.1-compat-ng-7.1.3` branch of material2 on a fork here
# since Angular and rules_typescript version under Bazel checking is too strict
# at the moment.
# https://github.com/gregmagolan/material2/commit/e2090864cddf926445eefd39c7e90eada107013d
# TODO(gregmagolan): update the next release of material that is compatible with
#   Angular 7.1.3 under Bazel
http_archive(
    name = "angular_material",
    sha256 = "75bec457885ddf084219a9da152ff79831d84909bb036552141ca3aadee64a04",
    strip_prefix = "material2-7.1.1-compat-ng-7.1.3",
    url = "https://github.com/gregmagolan/material2/archive/7.1.1-compat-ng-7.1.3.zip",
)

# Rules for compiling sass
http_archive(
    name = "io_bazel_rules_sass",
    sha256 = "76ae498b9a96fa029f026f8358ed44b93c934dde4691a798cb3a4137c307b7dc",
    strip_prefix = "rules_sass-1.15.1",
    url = "https://github.com/bazelbuild/rules_sass/archive/1.15.1.zip",
)

#######################################
# START DANS OWN DEPS
#######################################

# Rules for producing a GRPC gateway and translating protocol buffers to swagger definitions
http_archive(
    name = "io_grpc_grpc_java",
    sha256 = "0b86e44f9530fd61eb044b3c64c7579f21857ba96bcd9434046fd22891483a6d",
    strip_prefix = "grpc-java-1.18.0",
    url = "https://github.com/grpc/grpc-java/archive/v1.18.0.tar.gz",
)

grpc_gateway_version = "1.7.0"

http_archive(
    name = "grpc_ecosystem_grpc_gateway",
    strip_prefix = "grpc-gateway-{v}".format(v = grpc_gateway_version),
    url = "https://github.com/grpc-ecosystem/grpc-gateway/archive/v{v}.zip".format(v = grpc_gateway_version),
)

# Swagger Code Gen Jar for producing Angular HTTP Services
http_jar(
    name = "io_swagger_swagger_codegen_cli",
    url = "https://oss.sonatype.org/content/repositories/snapshots/io/swagger/swagger-codegen-cli/3.0.0-SNAPSHOT/swagger-codegen-cli-3.0.0-20180710.190537-87.jar",
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-bf564b4925ab5876a3f64d8b90fab7f769013d42",
    url = "https://github.com/bazelbuild/buildtools/archive/bf564b4925ab5876a3f64d8b90fab7f769013d42.zip",
)

#######################################
# END DANS OWN DEPS
#######################################

####################################
# Load and install our dependencies downloaded above.

load("@angular//packages/bazel:package.bzl", "rules_angular_dependencies")

rules_angular_dependencies()

load("@build_bazel_rules_typescript//:package.bzl", "rules_typescript_dependencies")

rules_typescript_dependencies()

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dependencies")

rules_nodejs_dependencies()

load("@build_bazel_rules_nodejs//:defs.bzl", "check_bazel_version", "node_repositories", "yarn_install")

# The minimum bazel version to use with this example repo
# Bazel 0.19 supports the .bazelignore file
check_bazel_version(minimum_bazel_version = "0.19.0")

node_repositories(
    node_version = "10.9.0",
    yarn_version = "1.12.1",
)

yarn_install(
    name = "npm",
    data = ["//:postinstall.tsconfig.json"],
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@io_grpc_grpc_java//:repositories.bzl", "grpc_java_repositories")

grpc_java_repositories()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@io_bazel_rules_webtesting//web:repositories.bzl", "browser_repositories", "web_test_repositories")

web_test_repositories()

browser_repositories(
    chromium = True,
    firefox = True,
)

load("@build_bazel_rules_typescript//:defs.bzl", "ts_setup_workspace")

ts_setup_workspace()

load("@io_bazel_rules_sass//:defs.bzl", "sass_repositories")

sass_repositories()

load("@angular//:index.bzl", "ng_setup_workspace")

ng_setup_workspace()

load("@angular_material//:index.bzl", "angular_material_setup_workspace")

angular_material_setup_workspace()

# Also define in Gopkg.toml
go_repository(
    name = "com_github_ghodss_yaml",
    commit = "0ca9ea5df5451ffdf184b4428c902747c2c11cd7",
    importpath = "github.com/ghodss/yaml",
)

# Also define in Gopkg.toml
go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
    importpath = "gopkg.in/yaml.v2",
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()