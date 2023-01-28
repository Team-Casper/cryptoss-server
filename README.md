# cryptoss-server

## API

| URL                                                           | Method | Description                                                    |
|---------------------------------------------------------------|--------|----------------------------------------------------------------|
| [/verification/start](#request-for-phone-authentication)      | `POST` | Request for phone authentication                               |
| [/verification/check](#authenticate-phone)                    | `POST` | Authenticate the phone using verification code                 |
| [/address](#register-account-address)                         | `POST` | Register account address under phone number                    |
| [/address/phone-number](#get-account-address-by-phone-number) | `GET`  | Get account address by phone number                            |
| [/send-coin](#send-asset-to-escrow)                           | `POST` | Send coin to unregisterd user. The coin will be sent to escrow |
| [/profile](#set-profile-picture-with-nft)                     | `POST` | Set profile picture with NFT                                   |


### Request for Phone Authentication

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

- Success Response
  - Code: 201
    - Content: `{"verification_code": "000000"}`


- Error Response
  - Code: 400
    - Content: `invalid phone number`

### Authenticate Phone

Authenticate the phone user has using the verification code.

- URL

`/verification/check`

- Method

`POST`

- Data Params

```json
{
  "phone-number": "01012345678",
  "verification-code": "123456"
}
```

- Success Response
  - Code: 200

- Error Response
  - Code: 401
    - Content: 
      - `invalid verification code`
      - `verification code expired`

### Register Account Address

Request to register account address under the phone number.

**[Advancement - not important for now]**

To prove that the user really owns the account and register the account under the phone number, a digital signature signed by the private key should also be included.
The message for signing would be like below:

```json
{
  "phone-number": "01012345678",
  "address": "0xa1b2c3d4..."
}
```

- URL

`/address`

- Method

`POST`

- Data Params

`sig` will not be implemented for now, but if possible, it would be good to add.

```json
{
  "phone-number": "01012345678",
  "address": "0xa1b2c3d4...",
  "sig": "<base64-encoded-signature>"
}
```

- Success Response
  - Code: 201

- Error Response
  - Code: 401
    - Content: `invalid siganture` (in advanced implementation)

### Get Account Address by Phone Number

Get account address by phone number.

- URL

`/address/:phone-number`

- Method

`GET`

- URL Params

`phone-number=[number]`

- Success Response
  - Code: 200
    - Content: `{ "account-address": "0xa1b2c3d4..."}`

- Error Response
  - Code: 400
    - Content: `invalid phone number`
  - Code: 404
    - Content: `not found under the phone number`

### Send Asset to Escrow

When sending an asset to a user who is not registered to cryptoss, the asset is sent to escrow.
When the receiver register to cryptoss, the asset will be sent to the receiver.
However, if the receiver do not register to cryptoss in x-days, the asset will be sent back to the sender.

- URL

`/send-coin`

- Method

`POST`

- Data Params

the unit of amount in `Octa` which is 10^-8 `APT`

```json
{
  "sender-phone-number": "01012345678",
  "receiver-phone-number": "01098765432",
  "amount": 100000000
}
```

- Success Response
  - Code: 201

- Error Response

### Set Profile Picture with NFT

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