# Steam Account Switcher
[<img src="https://raw.githubusercontent.com/Atrox/steam-account-switcher/master/icon/icon.ico" align="right" width="85">](https://github.com/atrox/steam-account-switcher/releases/latest)

A simpler and better way to switch between your steam accounts.
Everything works **without** touching your password.

![taskbar image](taskbar.png?raw=true "Taskbar Image")

## Installation & Setup

1. [Download the application](https://github.com/atrox/steam-account-switcher/releases/latest) and extract it somewhere you like
2. Update the file `accounts.toml` with your steam usernames in the following format
    ```toml
   [accounts] 
   username1234 = "some description"
   anotherusername = ""
   lastuser = "this is the last user"
   ```
3. Start the application and click on the icon in the taskbar whenever you want to switch accounts.

## Something does not work?

Look in the directory where the application lives.
If an error occurs the application creates/updates the file `error.log`.\
If the message does not help you, create an [issue](https://github.com/atrox/steam-account-switcher/issues).

## How does it work behind the scenes

It's a simple wrapper around the registry.
If you switch between the accounts in the registry, Steam remembers all your accounts without a problem.

Take a look at the code, it's really simple.

## Credits

The glorious icon is from [Flaticon.com created by SmashIcons](https://www.flaticon.com/authors/smashicons)
