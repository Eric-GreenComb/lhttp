# client of lhttp - a powerful go websocket framework

## usage

lhttp_client = new Lhttp("ws://localhost:8081/chat");

lhttp_client.on_open = function(context){
}

lhttp_client.on_message = function(context){
    context.send("hello, there!");
}

lhttp_client.on_close = function(context){
}

## Context API overview

    //set lhttp command
    setCommand = function(str) {
    }

    getCommand = function() {
    }

    getHeader = function(str) {
    }

    addHeader = function(key, value) {
    }

    getBody = function() {
    }

    //send lhttp body
    send = function(body) {
    }

    //publish a message to a channel
    publish = function(channel, command, headers, body) {
    }

    //subscribe a channel
    subscribe = function(channel, command, headers, body) {
    }