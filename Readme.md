# HTTP3 (quic) protocol implementation of Server + Client


### How to Run:

1. Generate Certificates

    ```sudo chmod +x generate_key.sh``` \
    host.cert , host.key will be generated

2. Run Server: (from server folder)

    ``` go run main.go ```

    Test with curl
    ```
        curl -vk --http3 https://127.0.0.1/hello                 
        *   Trying 127.0.0.1:443...
        * Server certificate:
        *  subject: C=In; ST=State; L=Province; O=Internet; OU=IT; CN=localhost
        *  start date: Mar  1 09:42:52 2026 GMT
        *  expire date: Mar  1 09:42:52 2027 GMT
        *  issuer: C=In; ST=State; L=Province; O=Internet; OU=IT; CN=localhost
        *  SSL certificate verify result: self signed certificate (18), continuing anyway.
        * Connected to 127.0.0.1 (127.0.0.1) port 443
        * using HTTP/3
        * [HTTP/3] [0] OPENED stream for https://127.0.0.1/
        * [HTTP/3] [0] [:method: GET]
        * [HTTP/3] [0] [:scheme: https]
        * [HTTP/3] [0] [:authority: 127.0.0.1]
        * [HTTP/3] [0] [:path: /]
        * [HTTP/3] [0] [user-agent: curl/8.12.1]
        * [HTTP/3] [0] [accept: */*]
        > GET / HTTP/3
        > Host: 127.0.0.1
        > User-Agent: curl/8.12.1
        > Accept: */*
        > 
        * Request completely sent off
        < HTTP/3 200 
        < date: Sun, 01 Mar 2026 09:48:52 GMT
        < content-length: 40
        < content-type: text/html
        < 
        * Connection #0 to host 127.0.0.1 left intact
        <html><body> Hello Http/3!</body></html>%                                                          
    ```

3. Run Client

    ```go run main.go``` \
    Response:
    ``` 
        Protocol: HTTP/3.0
        Status: 200 OK
        Body: {"message":"HTTP/3.0"}
    ```