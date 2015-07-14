(function(){
    'use strict'

    var hprose = require('hprose');

    var client = hprose.Client.create('http://127.0.0.1:7080');

    var person = client.useService(['Say', 'Hello']);
    person.Say('world', function(){});

    person.Hello(function(content){
        console.log(content); 
    });
})();
