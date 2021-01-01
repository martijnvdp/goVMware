param($credential)
if !($user) { $credential = get-creential }
$env:GOVMOMI_URL = "https://vc.mvdp.eu/sdk"
$env:GOVMOMI_USERNAME = $($Credential.username)
$env:GOVMOMI_PASSWORD = $($Credential.GetNetworkCredential().Password)
$env:GOVMOMI_INSECURE = "true"