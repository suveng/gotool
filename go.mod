module github.com/suveng/gotool

go 1.16

require (
	util v0.0.0
)

replace (
	util v0.0.0 => ./core
)