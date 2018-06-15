#!/bin/bash
#
# Copyright 2016 The Bazel Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
USAGE='bazel-run.sh [<bazel option>...] <target> [ -- [<target option>]... ]'
DESCRIPTION='
  Builds and runs the command generated by "bazel run" in the calling
  terminal. The command is run as a grandchild of the current shell,
  not from the Bazel server. Therefore, the program will have a controlling terminal, and
  the Bazel lock is released before running the command.'

function usage() {
  echo "$USAGE" "$DESCRIPTION" >&2
}

function die() {
  echo "$1"
  exit 1
}

function cleanup() {
  rm "$runcmd"
}

[ $# -gt 0 ] || { usage; exit 1; }

runcmd="$(mktemp /tmp/bazel-run.XXXXXX)" || die "Could not create tmp file"
trap "cleanup" EXIT

bazel run --script_path="$runcmd" "$@" || exit $?
[ -x "$runcmd" ] || die "File $runcmd not executable"

"$runcmd"