+++
title = 'Using DuckDB in AWS Lambda'
slug = 'using-duckdb-in-aws-lambda'
date = 2026-02-05 10:52:59
draft = false
tags = ['serverless', 'python', 'aws', 'lambda']
+++

This is going to be short but I wanted to get this down because I did this once before and everything worked out great but
then when I tried it again with another lambda project, I had a lot of problems getting credentials setup.

First of all, DO NOT go hard coding anything in the lambda. Definitely, DO NOT hard code credential key and secret but do not
hard code the region either, unless there is a *very specific* need for this like you're running the lambda in one region but 
say accessing an s3 bucket that is setup in another region.

The most prominent thing you will find for setting up s3 credentials in DuckDB is:

```sql
CREATE OR REPLACE SECRET secret (TYPE s3, PROVIDER credential_chain);
```

This is great, there is nothing wrong with it. It's going use the existing credentials on the lambda via roles to create the credentials for DuckDB.

The problem is that no directories are writable in the lambda except for `/tmp` and DuckDB is going to want to write the credentials somewhere.

So, you will want to run these items first...

```sql
SET home_directory = '/tmp';
SET secret_directory='/tmp';
SET extension_directory='/tmp';
```

Where I tripped up today was that was running the DuckDB query like this...

```python
import duckdb

duckdb.execute(f"""
    SET home_directory = '/tmp';
    SET secret_directory='/tmp';
    SET extension_directory='/tmp';
    CREATE OR REPLACE SECRET secret (TYPE s3, PROVIDER credential_chain);
    """)
```

...and getting errors "Invalid Input Error: Changing Secret Manager settings after the secret manager is used is not allowed!"

I tried this and still got the error which seemed weird because it should not have created if it existed, right?

```sql
CREATE IF NOT EXISTS secret (TYPE s3, PROVIDER credential_chain);
```

I went back to the lambda where I had it working perfectly checking syntax, steps, order to see where I had gone wrong.

It took a bit but then I noticed I was creating a connection first instead of executing directly from duckdb in the one where it worked.

```python
import duckdb

conn = duckdb.connect()
conn.execute(f"""
    SET home_directory = '/tmp';
    SET secret_directory='/tmp';
    SET extension_directory='/tmp';
    CREATE OR REPLACE SECRET secret (TYPE s3, PROVIDER credential_chain);
    """)
```

And it worked perfectly every single time it ran with no errors about not being able to change the secret manager.

I am not sure I fully understand it enough to explain it other than I create the connection and setup the credentials instead of trying to do it outside of a connection.