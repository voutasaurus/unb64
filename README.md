# unb64
parse base64 (std, url, padded, raw)

# usage

To parse standard base64 with padding from clipboard and write resulting binary to standard output
```
$ pbpaste | unb64
```

To parse unpadded URL base64 string of der encoded object from clipboard into openssl:
```
$ pbpaste | unb64 -r -u | openssl asn1parse -inform der
```

### note
pbcopy and pbpaste are command line utilities on mac OSX, they can be installed on windows by installing pasteboard with chocolatey.
