#!/bin/bash
set -e
echo "export RSA_PUBLIC_KEY='$(cat resources/keys/public.key)'" >>  test
echo "export RSA_PRIVATE_KEY='$(cat resources/keys/private.key)'" >>  test

