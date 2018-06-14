def _swagger_gen_impl(ctx):
    args = ctx.actions.args()
    args.add(["-jar", ctx.file._codegen_cli.path])
    args.add("generate")
    args.add(["-i", ctx.file.spec.path])
    args.add(["-l", "typescript-angular"])
    args.add(["-o", "{dirname}/{rule_name}".format(
        dirname=ctx.file.spec.dirname,
        rule_name=ctx.attr.name
    )])
    args.add(["--additional-properties", ctx.attr.additional_properties])

    outputs = []
    for out in ctx.attr.outs:
        outputs.append(ctx.actions.declare_file("%s/%s" % (ctx.attr.name, out)))

    ctx.actions.run(
        executable = ctx.executable._java,
        inputs = ctx.files._jdk + [
            ctx.executable._java,
            ctx.file._codegen_cli,
            ctx.file.spec
        ],
        outputs = outputs,
        arguments = [args],
        progress_message="Running Swagger Codegen %s" % ctx.label,
    )

    return struct(files = depset(outputs))

swagger_gen = rule(
    attrs = {
        # openapi spec file
        "spec": attr.label(
            mandatory = True,
            allow_single_file = FileType([".json", ".yaml"])
        ),
        "language": attr.string(
            mandatory = True
        ),
        "additional_properties": attr.string(),
        "outs": attr.string_list(),
        "_jdk": attr.label(
            default = Label("//tools/defaults:jdk"),
            allow_files = True
        ),
        "_java": attr.label(
            executable = True,
            cfg = "host",
            default = Label("@bazel_tools//tools/jdk:java"),
            allow_single_file = True,
        ),
        "_codegen_cli": attr.label(
            cfg = "host",
            default = Label("@io_swagger_swagger_codegen_cli//jar"),
            allow_single_file = True,
        ),
    },
    implementation = _swagger_gen_impl,
)
