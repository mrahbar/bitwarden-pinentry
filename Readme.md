This repo contains a [pinentry](https://www.gnupg.org/related_software/pinentry/index.html) implementation using Bitwarden as passphrase store.
It works by using utilizing the assuan protocol for gpg interaction and bitwarden-cli for the passphrase lookup.
Currently it is expected that the passphrase is the login password of an item.  

### Dependencies
````bitwarden-pinentry```` need the [Bitwarden-CLI](https://help.bitwarden.com/article/cli) to be installed on the machine.

### Configuration
````bitwarden-pinentry```` needs a configuration file specifing the bitwarden-cli sessionID (may also provided via ```` BW_SESSION ````  environment variable) and itemID (may also provided via ```` BW_ITEMID ````  environment variable).
````
{
  "Session": "...",
  "ItemID": "..."
}
 ```` 
 

### Limitations
This is an early working version of ````bitwarden-pinentry```` so be warned your mileage may vary.  