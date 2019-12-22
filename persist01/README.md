# Go to Pro - Persistence 
> This document assumes Go version 1.13.

## Table of contents ##
* [package pg02](#pg02)

<a name="pg02"/>
## package pg02
- Conection to Postgres with  "database/sql"


## FAQ - What is the equivalent of a Java ArrayList<E> in Golang?

``` java
class Channel {
    public String name;
    public Channel(){}
}

ArrayList<Channel> channels = new ArrayList<Channel>();

```


``` go
type Channel struct {
    name string
}
var channels []Channel  // an empty list
var channels = []*Channel {} // a empty list
channels = append(channels, Channel{name:"some channel name"})


```