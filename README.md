# goerror

Package goerror implements functions to manipulate errors.

The traditional `error` handling metyhod in golangis given below
which applied recursively up the call stack results in error reports without context or debugging information.

```
if err != nil {
        return err
}
```
The package mainly contain two **functions**

> GetErrorInfo
> New

## GetErrorInfo
which accepts the error and found the detailed information about the error.
`GetErrorInfo` returns the error wrapped up with the `file name`, 
`line number` and `function name` in which error occured.

> func GetErrorInfo(err error) error

```
if err != nil {
        return goerror.GetErrorInfo(err)
}
```

## New
Which is used to create a new errors.
New help us to create custom user defined errors.

> func New(err string) error

```
err := goerror.New("New custom error")
```
## Contributing

We welcome pull requests, bug fixes and issue reports. With that said, the bar for adding new symbols to this package is intentionally set high. Before proposing a change, please discuss your change by raising an issue.

## License

BSD-2-Clause
