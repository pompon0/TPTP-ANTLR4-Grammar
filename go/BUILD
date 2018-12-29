load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "parser",
    srcs = [
        "tptp_v7_0_0_0_base_listener.go",
        "tptp_v7_0_0_0_lexer.go",
        "tptp_v7_0_0_0_listener.go",
        "tptp_v7_0_0_0_parser.go",
    ],
    importpath = "github.org/pompon0/TPTP-ANTLR4-Grammar/go",
    visibility = ["//visibility:public"],
    deps = ["@com_github_antlr_antlr4//runtime/Go/antlr:go_default_library"],
)
