param($user,$pass)
$env:GOVMOMI_URL="https://vc.mvdp.eu/sdk"
$env:GOVMOMI_USERNAME=$user
$env:GOVMOMI_PASSWORD=$pass
$env:GOVMOMI_INSECURE="true"