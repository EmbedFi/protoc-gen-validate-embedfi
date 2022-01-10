Building
--------

This process is a hack!

The output is built from the proto definitions in the API repo.

```
protoc --go_out=. -I../api/proto ../api/proto/embedfi/finance/v1/*.proto
mv github.com/embedfi/finance/*.go .
rm -r github.com
```

But that's not all!

The field 'Value' in ScaledAmount conflicts with the interface for the
`database/sql/driver` Valuer interface, so you will also need to... I'm sorry
for this, edit the generated proto file to rename the Val field to Value

But wait, there's still more!

The String() method is overridden, you need to delete the String() method in
the generated code.

