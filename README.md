### Usage

* Download latest release to Termux:
    ```
    curl -L https://github.com/erdemolcay/termux-ssh-scripts/releases/download/v1.0.0/termux-ssh-scripts_1.0.0_android_arm64 --output /data/data/com.termux/files/usr/bin/termux-ssh-scripts
    ```
* Make executable:
    ```
    chmod +x /data/data/com.termux/files/usr/bin/termux-ssh-scripts```
    ```
* Install:
    ```
    termux-ssh-scripts install --api-token <CLOUDFLARE API TOKEN> --zone-id <CLOUDFLARE ZONE ID>
  
    ```
* (Optional) Run manually w/o installation and auto-sync:
    ```
    termux-ssh-scripts update --api-token <CLOUDFLARE API TOKEN> --zone-id <CLOUDFLARE ZONE ID>
  
    ```
