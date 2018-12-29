antlr4 -o ./go -Dlanguage=Go tptp_v7_0_0_0.g4
rm ./go/*.tokens
rm ./go/*.interp
