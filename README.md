# migrate-helper

tidb project need to [restructure tests](Tracking issue for restructure tests), migrate its old gocheck tests to 
testify tests. there are so many files need migrate, we can use this helper to do some regular replacements, then
manually check and update.

## install

```go
go install github.com/feitian124/migrate-helper
```

## usage

```shell
migrate-helper -f path/to/my_gocheck_test.go
```

then lines like `c.Assert(err, IsNil)` will be relaced to `require.NoError(t, err)`, then you can review and continue.

## supported 

- Equals
- IsNil
- Greater
- Less

...

see tests for full list.
