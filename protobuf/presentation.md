## Protocol Buffers

 - binary format for data serialization
 - open source
 - C++, C#, Java, Python, Go
 - also JavaScript, C, Perl, PHP, R, Julia, Scala, Swift, Haskell, Rust, etc.

### notes
 - structured data
 - numeric data is better

## Model your data

```protobuf
message Link {
  enum HttpMethod {
    GET   = 0;
    POST  = 1;
    PUT   = 2;
  }

  string href = 1;
  repeated HttpMethod methods = 2;
  string rel = 3;
}
```

### notes
 - you need to compile your model into a library for encoding/decoding  
 - keys are replaced by numeric ids (1 byte, 2 bytes)  
 - simple DSL (very similar to JSON) -> message descriptors  
 - msg descriptors can be reused across langs, platforms, etc.  
 - supports messages, enums, default values, maps, etc.  
 - fields are strongly typed  
 - data is validated  
 - extensible - newer versions of the model can read older data  

## Use
 - communication between frontend and backend  
 - communication between services  
 - data at rest  

## Code

Get JSON from file:
```go
  notifsJSON := &pb.Notifications{}
  notifsProto := &pb.Notifications{}

  json.Unmarshal([]byte(strData), &notifsJSON)
  proto.Unmarshal(binData, notifsProto)

  j := json.Marshal(notifsJSON)
  ioutil.WriteFile(filename, []byte(string(j)), 0644)

  pb := proto.Marshal(notifsProto)
  ioutil.WriteFile(filename, pb, 0644)
```

### notes
 - not so good for public APIs (need to provide library, increases complexity for clients)  
 - better suited for larger payloads but also works well with smaller ones  

## Pros & Cons
### Pros
 - smaller and faster than JSON  
 - structure validation (when used for RPC)  

### Cons
 - needs some setup  
 - not human-readable  

### notes
 - you can't directly play around with it and experiment  
 - network transfer times, compression times  

## Comparison to JSON

 - 10,000 order notifications  

Speed:
```
read JSON          446 ms
write JSON         371 ms

read Protobuf       60 ms  (~7.5x faster)
write Protobuf      62 ms  (~6x faster)
```

Data Size:
```
JSON:     22MB (14MB when compacted)
Protobuf: 11MB

Compressed JSON:      889KB (0.23s to compress)
Compressed Protobuf:  822KB (0.18s to compress)
```

## Links:
  - https://developers.google.com/protocol-buffers/
  - https://github.com/dcodeIO/protobuf.js
