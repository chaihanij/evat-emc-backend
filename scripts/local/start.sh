#!/bin/bash
set -e
# RSA_PUBLIC_KEY=$(cat resources/keys/public.key) \
# RSA_PRIVATE_KEY=$(cat resources/keys/private.key) \
# export RSA_PUBLIC_KEY=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCSQxqT2rO/Ro1cZRDzfLJPVCpFqp8Lrt+vRhv4w/u7ddxmrkwXloqogtKNwzBpnhAwVjCGRAyLyrekP4hTPAXo1R2unR+1mkt3CKJt+rsvzCzBmMHDX28BchexSgCRVMSQA8sy6wpD6Y/ftezZ5d0naIkkG/gMW0wms38wC5tTnwIDAQAB
# export RSA_PRIVATE_KEY=MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAJJDGpPas79GjVxlEPN8sk9UKkWqnwuu369GG/jD+7t13GauTBeWiqiC0o3DMGmeEDBWMIZEDIvKt6Q/iFM8BejVHa6dH7WaS3cIom36uy/MLMGYwcNfbwFyF7FKAJFUxJADyzLrCkPpj9+17Nnl3SdoiSQb+AxbTCazfzALm1OfAgMBAAECgYBrXgErmncqqXLp6XMJGneWjlU9sONx4xxnARNViSDI2ttBMR/AjQ0aaHaCBFAMqDCAGzUhyuBe2h/23YCIgV5gC0gkGUBFs7zlseiVcUzYvxUkqp8eQYmomgqm+1DbFVdi/So1r0dsCUvl9GadsPYlVTBMeaF0QVBYhiX2W5NdMQJBAO92OqgtC4OuP0P0JMhPr5zXr68V2KdOD47i3+dChLpqO4BpJ84h9yIYz5Mtty65+pNE6/h+Tzd3nIg/hNiz/wcCQQCcXRIMj8el19AqIl9YOHWEIzzgt493i44smdjXSK/bmtZ/baDqBQ3k5bN8YUogdjYvoOrQ3Q+A5UwVUpPpAkipAkEAkjkc2mMfO9qkMUedSmA2eUzPr9dkQ82L/JGXbe0VhYunJ+OOdfYRpdWGs3Xf/P1+AH2+pi/jCZjVkfFPmtBW8wJBAI8jfBRXwrIE4CDSGmoONTXjdCVapFERRfN6WtCpT6M/GUDJwNa6DcAOrsIb5dnVnUD4lDzjQQAD6CRU1eYmGxECQQDCPwFeb29LT4w56F2EdJAqvpAJj/upJpdTRWKKe94D2MDRmvw7fqCjTRfzDHiksn51s4h76kZzXVFthsgZvBW9
RSA_PUBLIC_KEY="
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAgDH0VQ4cQkuklpq/Uhoe
uy9oh8jefPT2nrk+sD81bEl7BOYjlJmiMi82+HwYODH5RlimOlyzaj49AfsXWgrq
ZbOqQoro52YixLYL0Eu1ug+E0i2EuVnlIWsMt+A81Z5WprNGgXAgdaS7ca6k4mIf
ESLQ2+aPitw8ti2Fge8lhstTxx5AF+B+vAmZxdUwK0ves9XGISeL+unNamcc1Gaf
+I0i+TVgvPKdQbUPOGkRJT38LHqLZfhvsH0PfmA4AoUz7E3euU0nwhvWr5mQYPjc
b2vp8RbNhlsgPs6Ni+6921WLD6Y/yATgJ72YrbsrqZ7w7ZEgpnzHIjVO3TN3ED7W
dwIDAQAB
-----END PUBLIC KEY-----
" \
RSA_PRIVATE_KEY="
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCAMfRVDhxCS6SW
mr9SGh67L2iHyN589PaeuT6wPzVsSXsE5iOUmaIyLzb4fBg4MflGWKY6XLNqPj0B
+xdaCupls6pCiujnZiLEtgvQS7W6D4TSLYS5WeUhawy34DzVnlams0aBcCB1pLtx
rqTiYh8RItDb5o+K3Dy2LYWB7yWGy1PHHkAX4H68CZnF1TArS96z1cYhJ4v66c1q
ZxzUZp/4jSL5NWC88p1BtQ84aRElPfwseotl+G+wfQ9+YDgChTPsTd65TSfCG9av
mZBg+Nxva+nxFs2GWyA+zo2L7r3bVYsPpj/IBOAnvZituyupnvDtkSCmfMciNU7d
M3cQPtZ3AgMBAAECggEAWlFXoDkxxa9/hS//jTuj3SN4lPn7brQzsP44YXoXDHT0
vC3ccYyQTqUfKJBdi/IBewNoEgRPz5GL8AJUoHbCT+b2FK3Di4xVAJx11EoS/cSD
QoUZVnE0UogWTNIqaiKFUCtMiFNUOy5VdbGir3CMDMmnttZI2X9zMrL9td6kZMOJ
aOGFSpH0PxzuIDqh05Ze8A0VduSznHNHLHO+ADrT7hblR4v9q9OzGJTA7CzbpTKS
Me1/nL9Dn/SxNGXFs6yhaMLx9h1O+8/R464myvuktqbKL3/H+C7lddn0WsLOTh/d
sLEP147JERWqn1S1lI+Kjv98dDvWJis64mNOqE9iMQKBgQDIoIVtYD81nUTv4jPF
mvrI2lbfB7cJ1XsFsI9W1/lpWNBzHyTFfptsq75OHm/zw2DRysYzEH7fBY6lrpWv
MtOoENx81HvaaYaHGkUjnHqU94OY4m4c8DY/rbYNQM0fUVcXdocm2KWPCG/I3jh8
4zOlrd7nrOq1qYQoQRStCvpeCQKBgQCjk7GfQnFk+Bs7ZnIVbcTJrcL6qPcvC+KJ
m74zRAFdDWwWIFO974UTTIS8t1y7PFD38oT7ahyW7n5BZfdJuZvF69GpbrRX5t6Q
j33IgP5malfagYs/IeEjXHSH6ybkQRh+lX8we+yPynzbede3ltCESrOVfsn7Jlm7
hcUWK0SwfwKBgQCsvemkXzaNhzgQA4foIKTFw5kAt7fV88XVBEymzZamKh5wwOP8
WCKi0s0snfBuxk7a0+kLlqxmgnZsGYIwM2ciUkJGCP4K9MksEmZtXxH8qZDZVjzR
FdLfyUXy61SyHmsKfLepruf925nELfIdNdMGWEqQ18XPXIG51y9iPw2LYQKBgEgc
6IVpaUw2BOte8pR32/V9YSPYMYDQIILB4kv/gwpezHPEtKZbXbNwXGRGAd+Za6ij
hfTAhvITGh4Csc4SwZWzrK6hW2gVI7FfUPh/xaeo4io2Sgj/Cp3oOIdjJ7Yg7IeB
qJzbRSLvBbRMNsl891gbBqi1SJ4r/gspogVv9cxJAoGALe2Q2P0efcnyEZqxYSPy
dNBVpY40ZjXjIzgb6ZSFrZAEsHTdBHB78VLrunLKPLCxZ+1SYw7HWI2sC0O0RH7N
L1eEq87g2daQwh4BFfUi9D2l64LQ1S7XPJADKQkyWDRFA5RGpNfEnHjZW+e+Gm9r
IiGZkw/zYuPNaOc/+bWhNGQ=
-----END PRIVATE KEY-----
" \
GIN_MODE=debug \
BASE_URL="http://localhost:8080" \
DEBUG=True \
MONGODB_URL='mongodb://0.0.0.0:37017' \
MONGODB_NAME=evat-emc-db \
MONGODB_USER=evatuser \
MONGODB_PASS=bHQ4mpnbCw6dp446 \
MONGODB_REQUEST_TIMEOUT=60 \
ENCRYPT_KEY=b054eb59dcf46db5da45ade306d005a2473def4a51f0ee93371bd64e77ae4b20 \
JWT_TOKEN_LIFE=0 air
#/usr/local/bin/go run app/cmd/main.go
