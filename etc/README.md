# Service Manager Configurations

## systemd

The `systemd/system` directory contains `systemd` service file `deployment-agent@.service` for deployment agent also known as a unit file. This file should be copied to [Unit File Load Path](https://www.freedesktop.org/software/systemd/man/systemd.unit.html#Unit%20File%20Load%20Path) usually at `/etc/systemd/system/`

This will help manually start and stop the service as follows.

```sh
# For reloading service manager configurations
sudo systemctl daemon-reload

sudo systemctl start deployment-agent@username.service
sudo systemctl stop deployment-agent@username.service

# To start on boot
sudo systemctl enable deployment-agent@username.service
```

There is another file in `systemd/user` folder which helps in starting or stopping the agent at user level (without `sudo`). To achieve that copy the file from `systemd/user` to `~/.config/systemd/user`

```sh
# For reloading user level systemd daemon
systemctl --user daemon-reload

systemctl --user start deployment-agent.service
systemctl --user stop deployment-agent.service
```

> Note: There are some changes needed in that file before it can be copied, such as $USERNAME and path to executable file `deployment-agent`

## upstart

The `upstart` directory contains a upstart configuration file for upstart service manager. This file should be copied to `/etc/init/system/` this will start with boot. It can also be started or stopped using following command
```sh
sudo initctl start deployment-agent
sudo initctl stop deployment-agent
```
