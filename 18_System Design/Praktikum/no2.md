## Fundamental 2 (20)
Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`
Dengan tujuan yang sama, tuliskan dalam bentuk perintah:
1. Redis
    ``` 
    KEYS users:*
    ```
2. Neo4j
    ```
    MATCH (u:User)
    RETURN u
    ```
3. Cassandra
    ```
    SELECT * FROM users
    ```