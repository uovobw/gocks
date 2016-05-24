gocks
=====

Small single-destination golang SOCKS5 proxy using github.com/armon/go-socks5
It is very dirty and not safe, for example it requires passwords in the command line, use at your own risk

Usage
=====

    usage: gocks [<flags>]
    
    Flags:
      -h, --help                   Show context-sensitive help (also try --help-long
                                   and --help-man).
      -u, --username="admin"       SOCKS username
      -p, --password="password"    SOCKS password
      -b, --bind=127.0.0.1         Address to bind to
      -d, --destination=127.0.0.1  Remote address that can be reached
      -P, --port="9186"            Port to bind to


