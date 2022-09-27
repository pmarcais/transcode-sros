module github.com/pmarcais/transcode-sros

go 1.16

require (
	github.com/pmarcais/transcode-sros/transsros v0.0.0-20220325230553-204967b3dea2
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli/v2 v2.4.0
)

replace github.com/pmarcais/transcode-sros/transsros => ./transsros
