param($credential, $prefix = "GOVC", $vchost)
if (!$credential) { $credential = get-credential }
if ($prefix -eq "GOVC") {
    $env:GOVC_URL = "https://" + $vchost
    $env:GOVC_USERNAME = $($Credential.username)
    $env:GOVC_PASSWORD = $($Credential.GetNetworkCredential().Password)
    $env:GOVC_INSECURE = "true"
}
if ($prefix -eq "GOVMOMI") {
    $env:GOVMOMI_URL = "https://" + $vchost + "/sdk"
    $env:GOVMOMI_USERNAME = $($Credential.username)
    $env:GOVMOMI_PASSWORD = $($Credential.GetNetworkCredential().Password)
    $env:GOVMOMI_INSECURE = "true"
}