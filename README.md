# fetch

a http request lib with go

### how to use

POST

```
url := "http://example.com"

data := map[string]string{
    "go":   "golang",
    "java": "javalang",
    "rust": "rustlang",
}
body, _ := json.Marshal(data)

header := http.Header{}
header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4")

result, err := fetch.Cmd("post", url, body, header)
fmt.Println(result, err)
```

PUT

```
url := "http://example.com"

data := map[string]string{
    "go":   "golang",
    "java": "javalang",
    "rust": "rustlang",
}
body, _ := json.Marshal(data)

header := http.Header{}
header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4")

result, err := fetch.Cmd("put", url, body, header)
fmt.Println(result, err)
```

GET

```
url := "http://example.com"

result, err := fetch.Cmd("get", url)
fmt.Println(result, err)
```

DELETE

```
url := "http://example.com"

result, err := fetch.Cmd("delete", url)
fmt.Println(result, err)
```
