# cryptoss-server

## API

| URL                                                      | Method | Description                                                    |
|----------------------------------------------------------|--------|----------------------------------------------------------------|
| [/verification/start](#request-for-phone-authentication) | `POST` | Request for phone authentication                               |
| [/verification/check](#authenticate-phone)               | `POST` | Authenticate the phone using verification code                 |
| [/account/address](#register-account-address)            | `POST` | Register account address under the phone number                |
| [/account/{phone-number}](#get-account-by-phone-number)  | `GET`  | Get account by phone number                                    |
| [/escrow](#send-asset-to-escrow)                         | `POST` | Send coin to unregisterd user. The coin will be sent to escrow |
| [/profile](#set-profile-picture-with-nft)                | `POST` | Set profile picture with NFT                                   |
| [/reset](#reset-account)                                 | `POST` | reset account                                                  |


## Request for Phone Authentication

Request to authenticate phone number. If phone number is valid, a 6-digit verification code expired in 5 minutes will be sent via sms.
If you want to receive a newly generated verification code (whatever the reason is), you can request again, and then the previous code will be expired. 

- URL

`/verification/start`

- Method

`POST`

- Data Params

```json
{
  "nickname": "<nickname>",
  "phone_number": "01012345678",
  "teleco_code": "SKT"
}
```

The `teleco_code` can be one of `SKT`, `KT`, `LGU`, `SKTMVNO`, `KTMVNO`, `LGUMVNO`

- Success Response
  - Code: 201
    - Content: `{"verification_code": "000000"}`


- Error Response
  - Code: 400
    - Content: `invalid phone number`

## Authenticate Phone

Authenticate the phone user has using the verification code.

- URL

`/verification/check`

- Method

`POST`

- Data Params

```json
{
  "phone_number": "01012345678",
  "verification-code": "123456"
}
```

- Success Response
  - Code: 201

- Error Response
  - Code: 401
    - Content: 
      - `invalid verification code`
      - `verification code expired`

## Register Account Address

Request to register account address to cryptoss account.

- URL

`/account/address`

- Method

`POST`

- Data Params

```json
{
  "phone_number": "01012345678",
  "address": "0xa1b2c3d4..."
}
```

- Success Response
  - Code: 201

- Error Response
  - Code: 401
    - Content: `invalid phone number`
  - Code: 404
    - Content: `account not found by the phone number`

## Get Account by Phone Number

Get account by phone number.

- URL

`/account/:phone-number`

- Method

`GET`

- URL Params

`phone-number=[number]`

- Success Response
  - Code: 200
    - Content: `{ "nickname": "myNickname", "account-address": "0xa1b2c3d4..."}`

- Error Response
  - Code: 400
    - Content: `invalid phone number`
  - Code: 404
    - Content: `not found under the phone number`

## Send Asset to Escrow

When sending an asset to a user who is not registered to cryptoss, the asset is sent to escrow.
When the receiver register to cryptoss, the asset will be sent to the receiver.
However, if the receiver do not register to cryptoss in x-days, the asset will be sent back to the sender.

- URL

`/escrow`

- Method

`POST`

- Data Params

the unit of amount in `Octa` which is 10^-8 `APT`

```json
{
  "sender_phone-number": "01012345678",
  "receiver_phone_number": "01098765432",
  "amount": 100000000
}
```

- Success Response
  - Code: 201

- Error Response

## Set Profile Picture with NFT

Set profile picture with a NFT owned by user.

- URL

`/profile`

- Method

`POST`

- Data Params

TBD

- Success Response
  - Code: 201

- Error Response
  - Code: 401
    - Content: `unauthorized for NFT`

## Reset Account

Reset account by phone number

- URL

`/reset`

- Method

`POST`

- Data Params

```json
{
  "phone_number": "01012345678"
}
```

- Success Response
  - Code: 201

- Error Response