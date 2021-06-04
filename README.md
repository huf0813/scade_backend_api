# SCADE Backend API

## :book: Description

Cancer impacted about nearly <b>10 million deaths in the world (WHO) by 2020</b>. One of the most common types of cancer
was skin (non-melanoma) with 1.2 million cases. Skin cancer was common but not everyone was aware of this cancer. Most
people not doing self-diagnoses, they might not understand how to do it periodically. Because of the late action, the
death rate of these cases is still on top of the chart. Through technology, self diagnosis can be provided in an easier
way. This may help to prevent skin cancer and reduce the death rate.

## :star: Minimum Viable Product (MVP)

* User can do a diagnosis
* User can view history of diagnosis
* User can view hospitals
* User can search hospitals by City
* User can book a first aid of skin cancer to hospital
* User can change hospitals of booking
* User can view a skin cancer's articles

## :rocket: How To Run

1. The app is combined with docker, make sure your environment support docker too.
1. Launch the app with a simple command ```yaml docker-compose up --build -d```
1. Well done! You are ready to go :partying_face:

## :shamrock: References

1. <b>Architecture</b> <br> We are different, not using MVC anymore. Using Clean Arch by Uncle Bob Martin, we have a
   chance to create multiple data source with multiple delivery. This time, we provide blazing fast API using GO + JSON
   Web Token + Go ORM (GORM), created a secure endpoints without excuses. Anyway, Thx
   to [bxcodec](https://github.com/bxcodec/go-clean-arch) has explained this examples.
1. <b>Upstream</b> <br> Our mission is helping user to mitigate skin cancer immediately.
   With [Indonesia's hospitals API](https://dekontaminasi.com/api/id/covid19/hospitals)
   from [Ariya Hidayat](https://github.com/ariya), Scade just successfully implement data integration with hospitals
   with simple ETL pipeline.

## :gift: Contributing and Publication

1. We couldn't wait your contribution. Please report the bugs by the issues
1. If you want to send a code. Please send your pull request to us, we would review your code immediately.

## :package: API Docs

1. /articles/image/{file : string}
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : image
   ```
1. /articles
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : 
   { 
    "success":true,
    "message":"fetch articles successfully",
    "data": [
        {"ID":1,
        "CreatedAt":"2021-06-03T13:51:53.054Z",
        "UpdatedAt":"2021-06-03T13:51:53.054Z",
        "DeletedAt":null,
        "title":"What is Cancer?",
        "body":"Cancer is a disease",
        "thumbnail":"default.jpg",
        "article_language_id":1}
    ]
   }
   ```
1. /articles/{language : string}
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : 
   { 
    "success":true,
    "message":"fetch articles successfully",
    "data": [
        {"ID":1,
        "CreatedAt":"2021-06-03T13:51:53.054Z",
        "UpdatedAt":"2021-06-03T13:51:53.054Z",
        "DeletedAt":null,
        "title":"What is Cancer?",
        "body":"Cancer is a disease",
        "thumbnail":"default.jpg",
        "article_language_id":1}
    ]
   }
   ```
1. /articles/{language : string}/{id : integer}
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   { 
    "success":true,
    "message":"fetch articles successfully",
    "data": {
        "ID":1,
        "CreatedAt":"2021-06-03T13:51:53.054Z",
        "UpdatedAt":"2021-06-03T13:51:53.054Z",
        "DeletedAt":null,
        "title":"What is Cancer?",
        "body":"Cancer is a disease",
        "thumbnail":"default.jpg",
        "article_language_id":1
    }
   }
   ```
1. /article_languages
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "fetch data successfully",
      "data": [
         {
            "ID": 1,
            "CreatedAt": "2021-06-03T13:51:53.033Z",
            "UpdatedAt": "2021-06-03T13:51:53.033Z",
            "DeletedAt": null,
            "language": "english",
            "articles": null
         },
         {
            "ID": 2,
            "CreatedAt": "2021-06-03T13:51:53.046Z",
            "UpdatedAt": "2021-06-03T13:51:53.046Z",
            "DeletedAt": null,
            "language": "indonesia",
            "articles": null
         }
      ]
   }
   ```
1. /diagnoses
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "fetch history of diagnoses successfully",
      "data": [
        {
         "ID": 11,
         "CreatedAt": "2021-06-04T08:42:06.55Z",
         "UpdatedAt": "2021-06-04T08:42:06.55Z",
         "DeletedAt": null,
         "cancer_name": "melanoma",
         "cancer_image": "be278254-c510-11eb-a487-0242ac120003_ISIC_0034319.jpg",
         "position": "neck",
         "price": 10,
         "user_id": 1,
         "invoices": null
        }
      ]
   }
   ```
1. /diagnoses/:id
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "fetch history of diagnoses successfully",
      "data": {
         "ID": 11,
         "CreatedAt": "2021-06-04T08:42:06.55Z",
         "UpdatedAt": "2021-06-04T08:42:06.55Z",
         "DeletedAt": null,
         "cancer_name": "melanoma",
         "cancer_image": "be278254-c510-11eb-a487-0242ac120003_ISIC_0034319.jpg",
         "position": "neck",
         "price": 10,
         "user_id": 1,
         "invoices": null
      }
   }
   ```
1. /diagnoses/image/:file
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : image
   ```
1. /diagnoses/create
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : 
    - multipart-form
        - field : 
            - cancer_name : string
            - cancer_image : file/image
            - position : neck
   Response Body : 
   {
      "success": true,
      "message": "data created successfully",
      "data": 11
   }
   ```
1. /hospitals
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : 
   {
    "success":true,
    "message":"get hospitals successfully",
    "data": [ 
        {"ID":1,
        "CreatedAt":"2021-06-03T13:51:53.351Z",
        "UpdatedAt":"2021-06-03T13:51:53.351Z",
        "DeletedAt":null,
        "name":"RS UMUM DAERAH DR. ZAINOEL ABIDIN",
        "address":"JL. TGK DAUD BEUREUEH, NO. 108 B. ACEH",
        "phone":"(0651) 34565",
        "region":"BANDA ACEH",
        "province":"Aceh",
        "invoices":null}
    ]
   }
   ```
1. /hospitals/search
   ```text
   HTTP Method : GET
   Query Params : 
   - city : string
   Authorization : none
   Request Body : none
   Response Body : 
   {
    "success":true,
    "message":"get hospitals by city successfully",
    "data": [
        {"ID":3,
        "CreatedAt":"2021-06-03T13:51:53.351Z",
        "UpdatedAt":"2021-06-03T13:51:53.351Z",
        "DeletedAt":null,
        "name":"RSUP SANGLAH",
        "address":"JL. DIPONEGORO DENPASAR BALI",
        "phone":"(0361) 227912",
        "region":"DENPASAR",
        "province":"Bali",
        "invoices":null}
    ]
   }
   ```
1. /hospitals/{id : integer}
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : none
   Response Body : 
   {
    "success":true,
    "message":"get hospital by id successfully",
    "data": {
        "ID":3,
        "CreatedAt":"2021-06-03T13:51:53.351Z",
        "UpdatedAt":"2021-06-03T13:51:53.351Z",
        "DeletedAt":null,
        "name":"RSUP SANGLAH",
        "address":"JL. DIPONEGORO DENPASAR BALI",
        "phone":"(0361) 227912",
        "region":"DENPASAR",
        "province":"Bali",
        "invoices":null
    }
   }
   ```
1. /invoices
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "fetch invoices successfully",
      "data": [
         {
         "invoice_id": 8,
         "hospital_name": "RS UMUM DAERAH  DR. ZAINOEL ABIDIN",
         "hospital_address": "JL. TGK DAUD BEUREUEH, NO. 108 B. ACEH",
         "hospital_phone": "(0651) 34565",
         "hospital_city": "BANDA ACEH",
         "hospital_province": "Aceh",
         "cancer_name": "melanoma",
         "cancer_image": "be278254-c510-11eb-a487-0242ac120003_ISIC_0034319.jpg",
         "cancer_position": "neck",
         "invoice_created_at": "2021-06-04T08:53:14.886Z",
         "invoice_updated_at": "2021-06-04T08:53:14.886Z"
         }
      ]
   }
   ```
1. /invoices/{id : integer}
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "fetch invoice by id successfully",
      "data": {
         "invoice_id": 8,
         "hospital_name": "RS UMUM DAERAH  DR. ZAINOEL ABIDIN",
         "hospital_address": "JL. TGK DAUD BEUREUEH, NO. 108 B. ACEH",
         "hospital_phone": "(0651) 34565",
         "hospital_city": "BANDA ACEH",
         "hospital_province": "Aceh",
         "cancer_name": "melanoma",
         "cancer_image": "be278254-c510-11eb-a487-0242ac120003_ISIC_0034319.jpg",
         "cancer_position": "neck",
         "invoice_created_at": "2021-06-04T08:53:14.886Z",
         "invoice_updated_at": "2021-06-04T08:53:14.886Z"
      }
   }
   ```
1. /invoices/create
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : 
   - JSON
   {
       "hospital_id" : 1,
       "diagnose_id" : 11
   }
   Response Body : 
   {
      "success": true,
      "message": "invoice created successfully",
      "data": null
   }
   ```
1. /profile
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : Bearer TOKEN
   Request Body : none
   Response Body : 
   {
      "success": true,
      "message": "get user's profile successfully",
      "data": {
         "ID": 1,
         "CreatedAt": "2021-06-03T13:51:53.198Z",
         "UpdatedAt": "2021-06-03T13:51:53.198Z",
         "DeletedAt": null,
         "name": "Harun Ulum Fajar",
         "address": "Malang",
         "email": "harun@gmail.com",
         "phone": "081308130813",
         "password": "$2a$10$lLLQ5ajPuWqzGjdioCN.su.EaZuxv5kYBBrsrgvaEgwkebb7w3yFO",
         "diagnoses": null,
         "subscription": null
      }
   }
   ```
1. /auth/sign_in
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : 
   - JSON
   {
       "email" : "harun@gmail.com",
       "password" : "1234567890"
   }
   Response Body : 
   {
      "success": true,
      "message": "sign in successfully",
      "data": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhhcnVuQGdtYWlsLmNvbSIsInJvbGUiOiJ1c2VyIiwiZXhwIjoxNjIyODMxODMzfQ.jqiYmRXlgbGcy48NBYJQy9ZUstAdIAfNtvfVxGjhgTE"
   }
   ```
1. /auth/sign_up
   ```text
   HTTP Method : GET
   Query Params : none
   Authorization : none
   Request Body : 
   - JSON
   {
       "name" : "harun",
       "address" : "malang",
       "email" : "email@gmail.com",
       "phone" : "1212",
       "password" : "dfasdfasdf"
   }
   Response Body : 
   {
      "success": true,
      "message": "sign up successfully",
      "data": null
   }
   ```
   wdyt? super cool, isn't? Cheers!
   ![cheers](https://media.giphy.com/media/kv5fbxHVAEOjrHeCLk/giphy.gif)