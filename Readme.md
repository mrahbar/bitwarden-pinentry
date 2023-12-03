# bitwarden-pinentry
This repo contains a [pinentry](https://www.gnupg.org/related_software/pinentry/index.html) implementation using Bitwarden as passphrase store.
It works by utilizing the [assuan protocol](http://info2html.sourceforge.net/cgi-bin/info2html-demo/info2html?(pinentry)Protocol) for gpg interaction and bitwarden-cli for the passphrase lookup.
Currently it is expected that the passphrase is the login password of an item stored in Bitwarden.  

### Dependencies
````bitwarden-pinentry```` need the [Bitwarden-CLI](https://help.bitwarden.com/article/cli) to be installed on the machine.

### Configuration
````bitwarden-pinentry```` needs a configuration file specifying the bitwarden-cli sessionID (may also be provided via ```` BW_SESSION ````  environment variable) and itemID (may also be provided via ```` BW_ITEMID ````  environment variable).
An minimal example would look like this:
````
{
  "Session": "...", // SessionID retrieved from bw login
  "ItemID": "...", // ID of secret retreived from visiting https://vault.bitwarden.com/#/vault
  "LogPath": "...", // Since StdOut is used for pinetry callback this file is used for Logs
  "EnableLog" : true // Enable or disable logs
}
 ```` 

## Usage
Simply specify the gpg-agent parameter ````pinentry-program```` to point to your ````bitwarden-pinentry```` binary in your ````gpg-agent.conf```` file.
For more on gpg-agent options visit this [site](https://www.gnupg.org/documentation/manuals/gnupg/Agent-Options.html).

### Limitations
This is an early working version of ````bitwarden-pinentry```` so be warned your mileage may vary.  