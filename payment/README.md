# Payment Service

![Payment service (1)](https://user-images.githubusercontent.com/23628103/165280680-79f31f1a-a7d6-48cb-983e-694be4d28ae4.png)

### Steps to setup locally:

Using Docker

```sh
# Build docker Image
> docker build -t cdp-team4/payment .

# Run docker image
> docker run -p 8008:8008 -d --name payment cdp-team4/payment

```

Without Docker

```sh
# Download golang dependencies
> go mod download

# Start the CMD server
> go run cmd/main.go
```

Run kafka with topic - "payment"

Open the below url
http://localhost:8008/swagger/index.html

### Routes supported

`GET ["/pay/"]` : healthcheck route of the the service
response:

```json
{
  "message": "service is running"
}
```

---

`POST ["/pay/"]` : initiate the payment process that create a record in DB and returns Razorpay's payment link.
request:

```json
{
  "Amount": 1000,
  "Currency": "INR",
  "OrderID": "asdfkjalksd",
  "UserID": "1591097270",
  "Method": "upi",
  "Description": "Payment for policy no #23456",
  "Notes": {}
}
```

response:

```json
// 200 : success
{
  "amount": 1000,
  "currency": "INR",
  "expire_by": 1691097057,
  "reference_id": "TS1989",
  "description": "Payment for policy no #23456",
  "customer": {
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com"
  },
  "notify": {
    "sms": true,
    "email": true
  },
  "reminder_enable": true,
  "notes": {},
  "callback_method": "get",
  "short_url": "https://rzp.io/i/nxrHnLJ",
  "status": "cancelled",
  "user_id": "1591097270"
}


// 400: failed
{
    "message":"unable to initiate payment process"
}
```

---

`GET [/pay/user/:user_id]`: get payment records of particular user. It is going to return list of all the payments initiated by the particular user.

request:

```json
{
  "user_id": "hafkldsasd"
}
```

response:

```json
// 200: success
[
  {
    "id": "qwrsdk23lkjn",
    "amount": 1000,
    "currency": "INR",
    "expire_by": 1691097057,
    "reference_id": "TS1989",
    "description": "Payment for policy no #23456",
    "notify": {
      "sms": true,
      "email": true
    },
    "notes": {},
    "status": "cancelled"
  },
  {
        "id": "laasdk23lkjn",
    "amount": 100,
    "currency": "INR",
    "expire_by": 1691099057,
    "reference_id": "TS19899",
    "description": "Payment for policy no #23456",
    "notify": {
      "sms": true,
      "email": true
    },
    "notes": {},
    "status": "confirmed"
  }
]

// 400 : failed
{
    "message":"unable to fetch payment records"
}
```

---

`GET ["/pay/:id"]` : get payment record of particular id.

request:

```json
{
  "payment_id": "lashfls234mn"
}
```

response:

```json
// 200: success
{
  "id": "lashfls234mn",
  "amount": 100,
  "currency": "INR",
  "expire_by": 1691099057,
  "reference_id": "TS19899",
  "description": "Payment for policy no #23456",
  "notify": {
    "sms": true,
    "email": true
  },
  "notes": {},
  "status": "confirmed"
}

// 400 : failed
{
    "message":"unable to fetch payment record"
}
```

---

`POST ["/pay/paymentMethods"]`

request:

```json
{
  "UserId": "aklshfl23",
  "Method": "bitcoin",
  "Agree": "1",
  "Comment": ""
}
```

response:

```json
// 200 : success
{ "message": "payment method added" }

// 400: failed
{ "message": "unable to add payment method" }

```

---

`GET ["/pay/paymentMethods/:id"]` : get list of all the payment methods supported for particular user

request:

```json
{
  "userID": "askfjlaksdj234"
}
```

response:

```json
// 200 : success
{
   "paymentmMethods": ["credit-card","debit-card","upi","net-banking"]
}

// 400: failed
{
    "message":"unable to fetch payment methods"
}
```