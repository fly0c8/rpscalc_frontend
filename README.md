# svelte-rps
- frontend for rpscalc nats calculator
- connects via websockets

Start nats with

nats-server -c my.cfg

#my.cfg
websocket {
  port: 8081
  no_tls: true
  # or configure tls { } if you want wss://
}




