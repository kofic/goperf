# goperf
Go based Load Tester

output shoudl be

output = {
    url:  string,
    time: datetime,
    bytes: bytes,
    runes: runes,
    body: string,
    headers: string,
    status: int
    
    // Run parameters

    jsassets: [
        {
            url:  string,
            time: datetime,
            bytes: bytes,
            runes: runes,
        },
        {...}
    ],

    cssassets: [
        {
            url:  string,
            time: datetime,
            bytes: bytes,
            runes: runes,
        },
        {...}
    ],

    // An Array of imgAsset response
    imgassets: [
        {
            url:  string,
            time: datetime,
            bytes: bytes,
            runes: runes,
        },
        {...}
    }
}

Run unit tests
go test ./... -cover

Run Benchmar
go test -bench BenchmarkResources -cpu 3

-cpu 3 uses 3 cores..default (without -cpu is 8 cores)
