# unb64
parse base64 (std, url, padded, raw)

# usage

To parse unpadded URL base64 string of der encoded object from clipboard into openssl:
```
$ pbpaste | unb64 -r -u | openssl asn1parse -inform der
```
