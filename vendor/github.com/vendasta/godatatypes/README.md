# Go Data Types
Since Go does not support generics this provides a generic templates
for data types to be generated for a specific type. This reduces the need for data types
to be rewritten for a different type.

All datatypes are tested in terms of their generic value types. You do not need to re-test type-specific data structures.

## Table of Contents
- [Supported Data Types](#supported-data-types)
- [Creating a New Data Type](#creating-a-new-data-type)
- [Future Work](#future-work)

## Supported Data Types
- [Set](set/internal/README.md) - Provides a set structure for storing unique values
- [Sync Map](syncmap/internal/README.md) - Thread safe map. Avoids the necessity for type assertions that comes with Go 1.9's  [sync map implementation](https://godoc.org/golang.org/x/sync/syncmap).

## Creating a New Data Type
The data types are generated with [genny](https://github.com/cheekybits/genny)
which allows a template for the data type to be created. Typed version of the
generic can then be generated from that file.

Each data structure should be put into its own directory which contains an internal
directory. The internal directory contains the actual generic template, and the script
used to generate the code. The script to generate the types is subject to change because
they are just bash scripts right now, and have lots of room for improvement.

There is a good example in [genny's docs](https://github.com/cheekybits/genny#real-example)
on how to implement a generic and generate the code for it.

## Future Work
- Generalize the generation strategy so there don't need to be bash files.

