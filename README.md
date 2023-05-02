go-safe4path
============

Convert any string to that available as filename in Windows

- `filename := safe4path.ToSafe(original,'%')`
    - Characters in `<>"/\|?*.:` and control codes in `original` are converted like  `%nn` (`nn` is ascii code)
- `safe4path.FromSafe(filename,'%')`
    - Get the original string from filename
- The original string should be a valid utf8 string.
- It does NOT check if the name is one of `NUL`,`CON`,`AUX` and other device names.
- **This package is a prototype. Do not use for production yet as it may contain issues**

References
----------
- [The weird world of Windows file paths | Fileside](https://www.fileside.app/blog/2023-03-17_windows-file-paths/)
