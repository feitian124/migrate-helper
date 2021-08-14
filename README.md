# migrate-helper

`tidb` project need to [restructure tests](https://github.com/pingcap/tidb/issues/26022), migrating its old `gocheck` tests to 
`testify` tests. there are so many files need to migrate that I write this helper to do the regular replacements, then
manually check and update.

## install

```go
go install github.com/feitian124/migrate-helper
```

## usage

```shell
migrate-helper path/to/my_gocheck_test.go
```

then lines like `c.Assert(err, IsNil)` will be relaced to `require.NoError(t, err)`, then you can review and continue.

## supported 

- Equals
- IsNil
- Greater
- Less

...

see tests for full list.
