# README Assignment 1
This covers the first assignment in PROG2005 Cloud Technologies

## Table of Contents
[Summary](#summary)  
[Expected Input Values and Return Values from Each Endpoint](#expected-input-values-and-return-values-from-each-endpoint)  
&emsp;&emsp;&emsp; [Status Endpoint](#status-endpoint)  
&emsp;&emsp;&emsp; [Exchange Endpoint](#exchange-endpoint)  
&emsp;&emsp;&emsp; [Info Endpoint](#info-endpoint)  
[External Services](#external-services)  
[Deployment/Setup](#deploymentsetup)  

## Summary
<p>
    A rest service that is provided with 3 endpoints:
</p>
<table>
    <tr>
        <th> Endpoints </th>
        <th> Information </th>
    </tr>
    <tr>
        <td> countryinfo/v1/status </td>
        <td> Returns the status code of 2 external rest services, local rest service version and local service uptime in secconds for the local rest servce </td>
    </tr>
    <tr>
        <td> countryinfo/v1/exchange </td>
        <td> Returns exchange rates for neighboring counties, this also takes into consideration if a coutry has more than 2 base currencies. For example Namibia</td>
    </tr>
    <tr>
        <td> countryinfo/v1/info </td>
        <td> Returns information regarding a country </td>
    </tr>
</table>

## Expected input values and return values from each endpoint
### Status endpoint
<p>Input value</p>
<code> countryinfo/v1/status </code>

<p>Expected output value</p>

``` json
{
    "RestCountriesApi":"200 OK", 
    "CurrenciesApi":"200 OK",
    "Version":"v1",
    "Uptime":7
}
```
Expected status-code `200 OK`

### exchange endpoint
<p>Input value</p>

<code> countryinfo/v1/exchange/{2_letter_ISO_code} </code>

<p>Example <code>ountryinfo/v1/exchange/no</code></p>
<p>Expected output</p>

> **_NB:_** If there are multiple base currencies, the last currency which is registered, is the one displayed in the GET reply

``` json
{
  "Base-Currency": [
    "NOK"
  ],
  "Country": "Norway",
  "Exchange-Rates": {
    "EUR": 0.088547,
    "RUB": 8.062744,
    "SEK": 0.941628
  }
}
```
Expected status-code `200 OK`

### Info endpoint
<p>Input value</p>

<code> countryinfo/v1/info{2_letter_ISO_code} </code>

<p>Example <code> countryinfo/v1/info/no </code></p>
<p>Expected output</p>

``` json
{
    "Borders": [
        "FIN",
        "SWE",
        "RUS"
    ],
    "Capital": [
        "Oslo"
    ],
    "Continents": [
        "Europe"
    ],
    "Flag": "https://flagcdn.com/w320/no.png",
    "Languages": {
        "nno": "Norwegian Nynorsk",
        "nob": "Norwegian Bokmål",
        "smi": "Sami"
    },
    "Name": "Norway",
    "Population": 5379475,
    "area": 323802
}
```
Expected status-code `200 OK`

## External services
This application relies on 2 external endpoints:
<p>
Rest Countries <code> http://129.241.150.113:8080/v3.1/ </code> <br>
Currency API <code> http://129.241.150.113:9090/currency/ </code>
</p>

## Deployment/Setup
The service is deployed in render. If no enviorment variable is available, the default port is chosen (8080)
<br> <br>
To access the service you can go to <br>

- Status : [countryinfo/v1/status](https://cloud-assignment-1-u94a.onrender.com/countryinfo/v1/status)
- Exchange : [countryinfo/v1/exchange/no](https://cloud-assignment-1-u94a.onrender.com/countryinfo/v1/exchange/no)
- Info : [countryinfo/v1/info/no](https://cloud-assignment-1-u94a.onrender.com/countryinfo/v1/info/no)

<br> <br>
To self deploy you clone down the project and type

```commandline
cd assignment-1
go build init
go run .
```



<!--
However, both setup and use should be documented for in a Readme 
(use the repository markdown language) alongside your codebase 
(more details below).

This notably includes a Readme file, which should at least include a general description and 
purpose of the software package/service, how to use (by specification and/or examples), how to deploy it, 
and what resources it relies on.
-->