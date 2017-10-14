# cd-go

## Configuration

All configurations are stored in a file located at $CDGO_CONFIG environment
variable if environment variable is unavailable it looks for
`$HOME/.config/cd-go/config.json`

## Commands

### `add` Command

* Add a new repository
* Repository's Webhook to be called.
* Repository Unique token

### `serve` Command

* Start server to listen for Webhook


> Note: The project structure will change through development


## Future Tasks
- [x] Start Gin Server to listen for Webhook Call.
- [x] Generate and Validate Token
- [ ] Hostname to be used.
- [x] Pull Repository on Webhook Call.
- [x] Add Whitelisting of IPs for Webhook Call.
- [x] Run Hook Commands/Script on Webhook Call.
- [ ] Add test cases
