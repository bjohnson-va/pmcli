# Generating New SyncMap From Generic

To generate a new set type from the generic run: `./gen KEYTYPE VALUETYPE`

Example: `./gen string bool` will generate the file `string_bool_syncmap.go` in the syncmap directory which
contains the generated code for a syncmap with string keys and bool values.
