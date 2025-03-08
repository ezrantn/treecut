package main

import (
	"log/slog"

	"github.com/ezrantn/treecut"
)

// This is a code example of how to use this library. In the `examples` directory, I created a `data` folder
// to store our example files, in this case `.txt`, but you could store anything.
//
// Here we are not partitioning by size, which means we are using the default value,
// partitioning by file count.
func main() {
	config := treecut.PartitionConfig{
		SourceDir:  "examples/data",
		OutputDirs: []string{"examples/partition1", "examples/partition2"},
		BySize:     false,
	}

	if err := treecut.MakePartitions(config); err != nil {
		slog.Error(err.Error())
	}

	// Unlink Example:
	//
	outputDirs := []string{"examples/partition1", "examples/partition2"}
	if err := treecut.RemovePartitions(outputDirs); err != nil {
		slog.Error(err.Error())
	}
}
