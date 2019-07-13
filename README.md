# HTML5 Speedtest

No Flash, No Java, No Websocket, No Bullshit.

This is a very lightweight Speedtest implemented in Javascript, using XMLHttpRequest and Web Workers.

Forked from [adolfintel/speedtest](https://github.com/adolfintel/speedtest) to remove unnecessary features and implement Go backend for use at [bgp.wtf](http://bgp.wtf).

## Compatibility
All modern browsers are supported: IE11, latest Edge, latest Chrome, latest Firefox, latest Safari.  
Works with mobile versions too.

## Features
* Download
* Upload
* Ping
* Jitter

## Server requirements
* A machine that can run Go code.

## Installation

    go install github.com/q3k/speedtest
    ./speedtest

## Donate

Support the original developer:
[![Donate with Liberapay](https://liberapay.com/assets/widgets/donate.svg)](https://liberapay.com/fdossena/donate)  
[Donate with PayPal](https://www.paypal.me/sineisochronic)  

## License
Copyright (C) 2016-2019 Federico Dossena

Copyright (C) 2019 Serge Bazanski

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
